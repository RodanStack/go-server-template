package main

import "github.com/spf13/cobra"

func NewServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		Long:  `This command starts the server. It is the entry point for the server.`,
	}
}
