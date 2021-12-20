package cmd

import (
	"fmt"
	"testing"
)

func Test_SetValue(t *testing.T) {
	vault := getVaultikData()

	err := vault.setValue("twitter_api_key415", "someImp0rt4ntK3y123")

	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := vault.getValue("twitter_api_key415", false)

	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	fmt.Println(res)
}

func Test_SetValue_delete(t *testing.T) {
	vault := getVaultikData()
	_key := "test_delete_file"
	err := vault.setValue(_key, "someImp0rt4ntK3y123")

	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	_, err = vault.getValue("twitter_api_key415", false)

	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	err = deleteFile(_key)

	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}
