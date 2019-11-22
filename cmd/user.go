package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user <action> [parameters]",
	Short: "Short",
	Long:  "Long",
}

func init() {
	if _, ok := os.LookupEnv(usernameKey); !ok {
		fmt.Printf("%s is not set but is required to work.\n", usernameKey)
		os.Exit(1)
	}

	if _, ok := os.LookupEnv(passwordKey); !ok {
		fmt.Printf("%s is not set but is required to work.\n", passwordKey)
		os.Exit(1)
	}
}
