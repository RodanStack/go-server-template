package main

import (
	"go-server-template/internal/cmd"
	"log"
)

func executeRootCmd() error {
	rootCmd := cmd.NewRootCmd()
	return rootCmd.Execute()
}

func main() {
	if err := executeRootCmd(); err != nil {
		log.Fatal(err)
	}
}
