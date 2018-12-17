package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "order",
	Short: "Simple HTTP order serivce",
	Run:   runServer,
}

// Execute runs the cobra rootCommand.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	f := rootCmd.Flags()
	f.StringVarP(&address, "address", "a", address, "listening address")
	f.StringVarP(&endpoint, "endpoint", "e", endpoint, "endpoint for other services to reach order service")
	f.StringVarP(&logLevel, "log-level", "l", logLevel, "log level (debug, info, warn, error), empty or invalid values will fallback to default")
	f.StringVarP(&redis, "redis-address", "r", redis, "redis address to connect to")
}
