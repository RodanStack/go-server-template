package cmd

import (
	"go-server-template/internal/app"

	"github.com/spf13/cobra"
)

func NewServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		Long:  `This command starts the server. It is the entry point for the server.`,
		Run: func(_ *cobra.Command, _ []string) {
			app.Run()
		},
	}
}
