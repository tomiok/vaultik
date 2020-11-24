package cmd

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

//setValue key is the actual key to identify the entry.
func (v *vaultik) setValue(key, value string) error {

	return nil
}

//getValue key is the actual key to identify the entry, return the encrypted/encoded value and an error (nil if any
// problem occurred).
func (v *vaultik) getValue(key string) (string, error) {
	return "", nil
}
