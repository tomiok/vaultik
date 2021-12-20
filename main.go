package main

import (
	"github.com/tomiok/vaultik/cmd"
	"log"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cannot execute cmd %v", err)
	}
}
