package main

import (
	"go-server-template/internal/cmd"
	"log"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	return godotenv.Load()
}

func executeRootCmd() error {
	rootCmd := cmd.NewRootCmd()
	return rootCmd.Execute()
}

func main() {
	if err := loadEnv(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := executeRootCmd(); err != nil {
		log.Fatal(err)
	}
}
