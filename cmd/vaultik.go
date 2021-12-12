package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	errNotFound  = errors.New("key not found")
	errFileEmpty = errors.New("file is empty")

	vault *vaultik
)

type vaultik struct {
	encodingKey string // is the key to encrypt the entry (the value given).
}

func newVaultik(encodingKey string) {
	s, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	p := filepath.Join(s, filepath.Base(".vaultik"))
	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.Write([]byte(encodingKey))

	if err != nil {
		fmt.Println(err)
		return
	}

}

//setValue key is the actual key to identify the entry.
func (v *vaultik) setValue(key, value string) error {

	if vault == nil {
		return errors.New("vault is nil, please run vaultik init -k [key]")
	}

	//check if the file exists
	_, err := os.OpenFile(key, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	encrypted, err := encrypt(v.encodingKey, value)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(key, []byte(encrypted), 0755)

	if err != nil {
		return err
	}

	return nil
}

//getValue key is the actual key to identify the entry, return the encrypted/encoded value and an error (nil if any
// problem occurred).
func (v *vaultik) getValue(key string) (string, error) {

	// key-values
	value, err := v.read(key)

	if err != nil {
		return "", err
	}

	return value, nil

}

// read the entire file, returns the values decrypted. The key is the file name
func (v *vaultik) read(key string) (string, error) {
	_, err := os.OpenFile(key, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadFile(key)

	if err != nil {
		return "", err
	}

	if len(b) == 0 {
		return "", errNotFound
	}

	decrypted, err := decrypt(v.encodingKey, string(b))

	if err != nil {
		return "", err
	}

	return decrypted, nil
}

func openVaultik() {

}
