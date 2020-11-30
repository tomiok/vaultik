package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

const separator = "\t" // key/value separator in the encrypted file
var errNotFound = errors.New("key not found")
var errFileEmpty = errors.New("file is empty")

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

// split the key value by the \t
func split(s string) (key, value string) {
	if s != "" {

		ss := strings.Split(s, separator)
		key = ss[0]
		value = ss[1]
		return
	}
	return "", ""
}

//setValue key is the actual key to identify the entry.
func (v *vaultik) setValue(key, value string) error {
	//check if the file exists
	_, err := os.OpenFile(v.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	// figure out if the value is here
	val, err := v.getValue(key)

	// if the value is already set, cannot pursue, should use UPDATE command instead
	if val != "" || !errors.Is(err, errNotFound) {
		return errors.New("cannot set an existing value, use UPDATE command instead")
	}

	str := joinkv(key, value)

	bytes, err := ioutil.ReadFile(v.filename)

	if err != nil {
		return err
	}

	decrypted, err := decrypt(v.encodingKey, string(bytes))

	if err != nil {
		if !errors.Is(err, errFileEmpty) {
			return err
		}
	}

	str = decrypted + "\n" + str

	encrypted, err := encrypt(v.encodingKey, str)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(v.filename, []byte(encrypted), 0755)

	if err != nil {
		return err
	}

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
	_, err := os.OpenFile(v.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(v.filename)

	if err != nil {
		return nil, err
	}

	if len(b) == 0 {
		return nil, errNotFound
	}

	decrypted, err := decrypt(v.encodingKey, string(b))

	if err != nil {
		return nil, err
	}

	kv := strings.Split(decrypted, "\n")

	return kv, nil
}

// make only one strign with key and value, glue with \t
func joinkv(key, value string) string {
	return key + separator + value
}
