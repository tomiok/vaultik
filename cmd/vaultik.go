package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	errNotFound  = errors.New("key not found")
	errFileEmpty = errors.New("file is empty")
)

const (
	dirSecure = "secure"
)

type vaultik struct {
	encodingKey string // is the key to encrypt the entry (the value given).
}

func newVaultik(encodingKey string) {
	f, err := openVaultikInHomeDir()

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

func getVaultikData() *vaultik {
	f, err := openVaultikInHomeDir()

	if err != nil {
		log.Println(err)
		return nil
	}

	buf := make([]byte, 128)
	_, err = f.Read(buf)

	return &vaultik{encodingKey: string(buf)}
}

//setValue key is the actual key to identify the entry.
func (v *vaultik) setValue(key, value string) error {
	// get user home dir
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	err = os.Mkdir(filepath.Join(home, dirSecure), os.ModePerm)

	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	p := filepath.Join(home, dirSecure, filepath.Base(key))

	fmt.Println(fmt.Sprintf("directory created: %s", p))
	// check if the file exists. But not append anything, just overwrite.
	varFile, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY, os.ModePerm)

	if err != nil {
		return err
	}

	//encrypt the value with the encoding key
	encrypted, err := encrypt(v.encodingKey, value)

	if err != nil {
		return err
	}

	_, err = varFile.Write([]byte(encrypted))

	if err != nil {
		return err
	}

	return nil
}

//getValue key is the actual key to identify the entry, return the encrypted/encoded value and an error (nil if any
// problem occurred).
func (v *vaultik) getValue(key string, decrypted bool) (string, error) {

	// key-values
	value, err := v.read(key, decrypted)

	if err != nil {
		return "", err
	}

	return value, nil

}

// read the entire file, returns the values decrypted. The key is the file name
func (v *vaultik) read(key string, decrypted bool) (string, error) {
	// get user home dir
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	p := filepath.Join(home, dirSecure, filepath.Base(key))

	res, err := os.ReadFile(p)

	if err != nil {
		return "", err
	}

	if len(res) == 0 {
		return "", errNotFound
	}

	if decrypted {
		d, err := decrypt(v.encodingKey, string(res))

		if err != nil {
			return "", err
		}

		return d, nil
	}

	return string(res), nil
}

func openVaultikInHomeDir() (*os.File, error) {
	s, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	p := filepath.Join(s, filepath.Base(".vaultik"))

	return os.OpenFile(p, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}
