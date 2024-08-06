package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rootCmd := NewRootCmd()
	if err = rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing App: %v", err)
	}
}
