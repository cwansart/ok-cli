package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user <action> [parameters]",
	Short: "Manage users",
	Long:  "Long",
}

func init() {
	if _, ok := os.LookupEnv(userNameKey); !ok {
		fmt.Printf("%s is not set but is required to work.\n", userNameKey)
		os.Exit(1)
	}

	if _, ok := os.LookupEnv(passwordKey); !ok {
		fmt.Printf("%s is not set but is required to work.\n", passwordKey)
		os.Exit(1)
	}
}
