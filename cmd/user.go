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
	if _, ok := os.LookupEnv(userNameKey); !ok {
		fmt.Println(userNameKey + " is not set but is required to work.")
		os.Exit(1)
	}

	if _, ok = os.LookupEnv(passwordKey); !ok {
		fmt.Println(passwordKey + " is not set but is required to work.")
		os.Exit(1)
	}
}
