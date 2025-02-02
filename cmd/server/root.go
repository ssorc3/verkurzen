package main

import (
    "github.com/spf13/cobra"
)

var (
    rootCmd = &cobra.Command{
        Use: "server",
        Short: "server is example cmd app with golang",
        Long: `server is example golang web service using gin and cassandra`,
    }
)

func Execute() error {
    return rootCmd.Execute()
}
