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
	Use:     "migrate {user} {path to id_rsa/password} {destination} {remoteURL} [-m=privateKey]",
	Short:   "migrate should send in scp protocol all the secured and hashed pair of key-values",
	Long:    ``,
	Example: "migrate tomas [p45w0rd1 || ] /home/secure example.com:22",
	Args:    cobra.MinimumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {

		username := args[0]
		access := args[1]
		destination := args[2]
		host := args[3]
		///////////////////////////// TODO contar y mostrar CUANTOS files se migraron
		var client *ssh.Client
		var err error
		if method == "" {
			cfg := withPassword(username, access, nil)
			client, err = ssh.Dial("tcp", host, &cfg)
		} else {
			cfg, err := withPrivateKey(username, access, nil)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			client, err = ssh.Dial("tcp", host, &cfg)
		}

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		count := walk(destination, client)

		fmt.Println(fmt.Sprintf("%d files migrated", count))
	},
}

func init() {
	migrateCmd.Flags().StringVarP(&method, "method", "m", "", "method is the auth method for"+
		"ths SSH connection. use privateKey and provide the PATH to your private key, or leave it blank to use "+
		"standard user + password connection")
	rootCmd.AddCommand(migrateCmd)
}

func walk(destination string, client *ssh.Client) int {
	var count int
	p, err := getSecretDirectory()

	if err != nil {
		fmt.Println(fmt.Sprintf("cannot get secret dir dir %s", err.Error()))
		return count
	}

	f, err := os.Open(p)

	if err != nil {
		fmt.Println(fmt.Sprintf("cannot get secret dir dir %s", err.Error()))
		return count
	}

	files, err := f.ReadDir(0)

	if err != nil {
		fmt.Println(fmt.Sprintf("cannot get secret dir dir %s", err.Error()))
		return count
	}
	var wg sync.WaitGroup

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		sess, _ := client.NewSession()
		wg.Add(1)
		go func(name string) {
			err = doCopy(&wg, p, name, filepath.Join(destination, name), sess)

			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(fmt.Sprintf("migrating secure file %s ", name))
		}(file.Name())
		count++
	}
	wg.Wait()
	return count
}

func doCopy(wg *sync.WaitGroup, p, name, destination string, sess *ssh.Session) error {
	defer wg.Done()
	join := filepath.Join(p, name)
	return copyPath(join, destination, sess)
}

func copyPath(filePath, destinationPath string, session *ssh.Session) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	s, err := f.Stat()

	if err != nil {
		return err
	}

	return copyFile(s.Size(), s.Mode().Perm(), path.Base(filePath), f, destinationPath, session)
}

func copyFile(size int64, mode os.FileMode, fileName string, contents io.Reader, destination string, session *ssh.Session) error {
	defer func() {
		_ = session.Close()
	}()

	w, err := session.StdinPipe()

	if err != nil {
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

	_, _ = fmt.Fprintf(w, "C%#o %d %s\n", mode, size, fileName)
	_, _ = io.Copy(w, contents)
	_, _ = fmt.Fprint(w, "\x00")
	_ = w.Close()

	return <-errors
}
