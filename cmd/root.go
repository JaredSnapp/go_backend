package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "go_backend",
	Short: "Go backend dev",
}

func Execute() {
	RootCmd.AddCommand(serviceCmd)

	// conf := config.Get()

	if err := RootCmd.Execute(); err != nil {
		log.Fatal("Failed to start service", err)
	}
}
