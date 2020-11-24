package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

const separator = "\t" // key/value separator in the encrypted file
var errNotFound = errors.New("key not found")

type vaultik struct {
	encodingKey string // is the key to encrypt the entry (the value given).
	filename    string // where the entry is going to be stored.
}

func newVaultik(encodingKey, filename string) vaultik {
	return vaultik{
		encodingKey: encodingKey,
		filename:    filename,
	}
}

func split(s string) (key, value string) {
	ss := strings.Split(s, separator)
	key = ss[0]
	value = ss[1]
	return
}

//setValue key is the actual key to identify the entry.
func (v *vaultik) setValue(key, value string) error {

	return nil
}

//getValue key is the actual key to identify the entry, return the encrypted/encoded value and an error (nil if any
// problem occurred).
func (v *vaultik) getValue(key string) (string, error) {

	// key-values
	kv, err := v.readAll()

	if err != nil {
		return "", err
	}

	for _, entry := range kv {
		k, v := split(entry)

		if k == key {
			return v, nil
		}
	}

	return "", errNotFound

}

// readAll reads the entire file, returns the values decrypted
func (v *vaultik) readAll() ([]string, error) {
	if _, err := os.Stat(filename); err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	decrypted, err := decrypt(v.encodingKey, string(b))

	if err != nil {
		return nil, err
	}

	kv := strings.Split(decrypted, "\b")

	return kv, nil
}
