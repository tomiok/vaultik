package cmd

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

// withPrivateKey connect ssh session with user and password
func withPrivateKey(username string, path string, keyCallBack ssh.HostKeyCallback) (ssh.ClientConfig, error) {
	if keyCallBack == nil {
		keyCallBack = ssh.InsecureIgnoreHostKey()
	}

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

// withPassword connect ssh session with user and password
func withPassword(username string, password string, keyCallBack ssh.HostKeyCallback) ssh.ClientConfig {
	if keyCallBack == nil {
		keyCallBack = ssh.InsecureIgnoreHostKey()
	}

	return ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: keyCallBack,
	}
}

// withPrivateKeyWithPassphrase Creates the configuration for a client that authenticates with a password protected private key
func withPrivateKeyWithPassphrase(username string, passPhrase []byte, path string, keyCallBack ssh.HostKeyCallback) (ssh.ClientConfig, error) {
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
