package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:              "go-server-template",
		Short:            "Commander for go-server-template",
		Long:             `This is the root command for go-server-template. It is the entry point for all other commands.`,
		TraverseChildren: true,
	}

	rootCmd.AddCommand(NewServeCmd())
	return rootCmd
}
