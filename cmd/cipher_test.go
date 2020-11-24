package cmd

import (
	"fmt"
	"testing"
)

var key = "sophisticated_key"

func Test_encrypt(t *testing.T) {
	value := "one_time_value"
	s, err := encrypt(key, value)

	if err != nil {
		t.Fail()
	}

	if s == "" {
		t.Fail()
	}

	fmt.Print(s)
}

func Test_decrypt(t *testing.T) {
	// one_time_value
	cipher := "fa1a2b0967d8bd97be6cf52cceee0a924dbe63fcc2435987e2371cc011f5"
	s, err := decrypt(key, cipher)

	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if s != "one_time_value" {
		t.Fail()
	}
}
