package cmd

import (
	"fmt"
	"github.com/kballard/go-shellquote"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path"
	"path/filepath"
	"sync"
)

// method is for
var method string

var migrateCmd = &cobra.Command{
	Use:     "migrate {user} {path to id_rsa} {destination} [-m=privateKey]",
	Short:   "migrate should send in scp protocol all the secured and hashed pair of key-values",
	Long:    ``,
	Example: "migrate tomas p45w0rd1 example.com:22",
	Args:    cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		username := args[0]
		access := args[1]
		host := args[2]

		if method == "" {

		} else {
			cfg, _ := withPrivateKey(username, access, nil)
			client, _ := ssh.Dial("tcp", host, &cfg)
			walk(client)
		}
	},
}

func init() {
	migrateCmd.Flags().StringVarP(&method, "method", "m", "", "method is the auth method for"+
		"ths SSH connection. use privateKey and provide the PATH to your private key, or leave it blank to use "+
		"standard user + password connection")
	rootCmd.AddCommand(migrateCmd)
}

func walk(client *ssh.Client) {
	p, err := getSecretDirectory()

	if err != nil {
		fmt.Println(fmt.Sprintf("cannot get secret dir dir %s", err.Error()))
		return
	}

	f, err := os.Open(p)

	if err != nil {
		fmt.Println(fmt.Sprintf("cannot get secret dir dir %s", err.Error()))
		return
	}

	files, err := f.ReadDir(0)

	if err != nil {
		fmt.Println(fmt.Sprintf("cannot get secret dir dir %s", err.Error()))
		return
	}
	var wg sync.WaitGroup

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		sess, _ := client.NewSession()
		wg.Add(1)
		go func(name string) {
			doCopy(&wg, p, name, name, sess)
		}(file.Name())
	}
	wg.Wait()
}

func doCopy(wg *sync.WaitGroup, p, name, destination string, sess *ssh.Session) {
	defer wg.Done()
	join := filepath.Join(p, name)
	CopyPath(join, destination, sess)
}

func CopyPath(filePath, destinationPath string, session *ssh.Session) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	s, err := f.Stat()
	if err != nil {
		return err
	}
	return copyFile(s.Size(), s.Mode().Perm(), path.Base(filePath), f, destinationPath, session)
}

func copyFile(size int64, mode os.FileMode, fileName string, contents io.Reader, destination string, session *ssh.Session) error {
	defer session.Close()
	w, err := session.StdinPipe()

	if err != nil {
		fmt.Println("err2: " + err.Error())
		return err
	}

	cmd := shellquote.Join("scp", "-t", destination)
	if err = session.Start(cmd); err != nil {
		_ = w.Close()
		return err
	}

		errors := make(chan error)

	go func() {
			errors <- session.Wait()
	}()

	fmt.Fprintf(w, "C%#o %d %s\n", mode, size, fileName)
	io.Copy(w, contents)
	fmt.Fprint(w, "\x00")
	w.Close()

	return <-errors
}
