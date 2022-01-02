package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

type Client struct {
	// the host to connect to
	Host string

	// the client config to use
	ClientConfig *ssh.ClientConfig

	// stores the SSH session while the connection is running
	Session *ssh.Session

	// stores the SSH connection itself in order to close it after transfer
	Conn ssh.Conn

	// the maximal amount of time to wait for a file transfer to complete
	Timeout time.Duration

	// the absolute path to the remote SCP binary
	RemoteBinary string
}

// Connect to the remote SSH server, returns error if it couldn't establish a session to the SSH server
func (a *Client) Connect() error {
	if a.Session != nil {
		return nil
	}

	client, err := ssh.Dial("tcp", a.Host, a.ClientConfig)
	if err != nil {
		return err
	}

	a.Conn = client.Conn
	a.Session, err = client.NewSession()
	if err != nil {
		return err
	}
	return nil
}

// CopyFromFile Copies the contents of an os.File to a remote location, it will get the length of the file by looking it up from the filesystem
func (a *Client) CopyFromFile(file os.File, remotePath string, permissions string) error {
	info, _ := file.Stat()
	return a.Copy(&file, remotePath, permissions, info.Size())
}

// Copy the contents of an io.Reader to a remote location.
// Access copied bytes by providing a PassThru reader factory
func (a *Client) Copy(r io.Reader, remotePath string, permissions string, size int64) error {
	stdout, err := a.Session.StdoutPipe()
	ctx := context.Background()
	if err != nil {
		return err
	}

	filename := path.Base(remotePath)

	wg := sync.WaitGroup{}
	wg.Add(2)

	errCh := make(chan error, 2)

	go func() {
		defer wg.Done()
		w, err := a.Session.StdinPipe()
		if err != nil {
			errCh <- err
			return
		}

		defer w.Close()

		_, err = fmt.Fprintln(w, "C"+permissions, size, filename)
		if err != nil {
			errCh <- err
			return
		}

		if err = checkResponse(stdout); err != nil {
			errCh <- err
			return
		}

		_, err = io.Copy(w, r)
		if err != nil {
			errCh <- err
			return
		}

		_, err = fmt.Fprint(w, "\x00")
		if err != nil {
			errCh <- err
			return
		}

		if err = checkResponse(stdout); err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		err := a.Session.Run(fmt.Sprintf("%s -qt %q", a.RemoteBinary, remotePath))
		if err != nil {
			errCh <- err
			return
		}
	}()

	if a.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, a.Timeout)
		defer cancel()
	}

	if err := wait(&wg, ctx); err != nil {
		return err
	}

	close(errCh)
	for err := range errCh {
		if err != nil {
			return err
		}
	}
	return nil
}

func wait(wg *sync.WaitGroup, ctx context.Context) error {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()

	select {
	case <-c:
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

// checkResponse check the response it reads from the remote, and will return a single error in case
// of failure
func checkResponse(r io.Reader) error {
	response, err := ParseResponse(r)
	if err != nil {
		return err
	}

	if response.IsFailure() {
		return errors.New(response.GetMessage())
	}

	return nil

}

type ResponseType = uint8

const (
	Ok      ResponseType = 0
	Warning ResponseType = 1
	Error   ResponseType = 2
)

type Response struct {
	Type    ResponseType
	Message string
}

// ParseResponse Reads from the given reader (assuming it is the output of the remote) and parses it into a Response structure
func ParseResponse(reader io.Reader) (Response, error) {
	buffer := make([]uint8, 1)
	_, err := reader.Read(buffer)
	if err != nil {
		return Response{}, err
	}

	resType := buffer[0]
	message := ""
	if resType > 0 {
		bufferedReader := bufio.NewReader(reader)
		message, err = bufferedReader.ReadString('\n')
		if err != nil {
			return Response{}, err
		}
	}

	return Response{resType, message}, nil
}

func (r *Response) IsOk() bool {
	return r.Type == Ok
}

func (r *Response) IsWarning() bool {
	return r.Type == Warning
}

// IsError returns true when the remote responded with an error
func (r *Response) IsError() bool {
	return r.Type == Error
}

// IsFailure returns true when the remote answered with a warning or an error
func (r *Response) IsFailure() bool {
	return r.Type > 0
}

// GetMessage returns the message the remote sent back
func (r *Response) GetMessage() string {
	return r.Message
}

///////////////////
// client config //
///////////////////

func NewClient(host string, config *ssh.ClientConfig) Client {
	return NewConfig(host, config).Create()
}

//ClientConfig is the set of variables for the ssh client
type ClientConfig struct {
	host         string
	clientConfig *ssh.ClientConfig
	session      *ssh.Session
	timeout      time.Duration
	remoteBinary string
}

// NewConfig Creates a new client config.
// It takes the required parameters: the host and the ssh.ClientConfig and
// returns a config populated with the default values for the optional
// parameters.
//
// These optional parameters can be set by using the methods provided on the
// ClientConfig struct.
func NewConfig(host string, config *ssh.ClientConfig) *ClientConfig {
	return &ClientConfig{
		host:         host,
		clientConfig: config,
		timeout:      2 * time.Minute, // no timeout by default
		remoteBinary: "scp",
	}
}

// Create Builds a client with the configuration stored within the ClientConfigurer
func (c *ClientConfig) Create() Client {
	return Client{
		Host:         c.host,
		ClientConfig: c.clientConfig,
		Timeout:      c.timeout,
		RemoteBinary: c.remoteBinary,
		Session:      c.session,
	}
}
