package cmd

import (
	"os"
	"path/filepath"
)

// getSecretPath use it to get the file location, the key is also the file name
func getSecretPath(key string) (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	p := filepath.Join(home, dirSecure, key)

	return p, nil
}

func createSecretDirectory() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	return os.Mkdir(filepath.Join(home, dirSecure), os.ModePerm)
}

func getSecretDirectory() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	p := filepath.Join(home, dirSecure)

	return p, nil
}

func getOrCreateSecureFile(key string) (*os.File, error) {
	p, err := getSecretPath(key)
	if err != nil {
		return nil, err
	}

	// check if the file exists. But not append anything, just overwrite.
	return os.OpenFile(p, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}

func readSecureFile(key string) ([]byte, error) {
	p, err := getSecretPath(key)

	if err != nil {
		return nil, err
	}

	return os.ReadFile(p)
}
