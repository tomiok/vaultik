package cmd

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

// PrivateKey Loads a private and public key from "path" and returns an SSH ClientConfig to authenticate with the server
// i.e. client `Config, _ := auth.PrivateKey("username", "/path/to/rsa/key", ssh.InsecureIgnoreHostKey())`
func PrivateKey(username string, path string, keyCallBack ssh.HostKeyCallback) (ssh.ClientConfig, error) {
	privateKey, err := ioutil.ReadFile(path)

	if err != nil {
		return ssh.ClientConfig{}, err
	}

	signer, err := ssh.ParsePrivateKey(privateKey)

	if err != nil {
		return ssh.ClientConfig{}, err
	}

	return ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: keyCallBack,
	}, nil
}

// PrivateKeyWithPassphrase Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(username string, passPhrase []byte, path string, keyCallBack ssh.HostKeyCallback) (ssh.ClientConfig, error) {
	privateKey, err := ioutil.ReadFile(path)

	if err != nil {
		return ssh.ClientConfig{}, err
	}
	signer, err := ssh.ParsePrivateKeyWithPassphrase(privateKey, passPhrase)

	if err != nil {
		return ssh.ClientConfig{}, err
	}

	return ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: keyCallBack,
	}, nil
}
