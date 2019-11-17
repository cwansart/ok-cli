package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "", // TODO: add descriptions
	Long:  "",
	Run:   userList,
}

func userList(_ *cobra.Command, _ []string) {
	// TODO: check args if there is a user name and access that instead. Do we need another sub command for that?
	req, err := http.NewRequest("GET", userListUrl(), nil)

	if err != nil {
		fmt.Println("An error occured during request creation: ", err) // TODO: use logger or an error output instead?
		return
	}

	req.SetBasicAuth(os.Getenv(userNameKey), os.Getenv(passwordKey))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("An error occured during request: ", err)
		return
	}

	respBody, _ := ioutil.ReadAll(resp.Body) // TODO: error handling

	// TODO: pretify JSON output or should we just output the names?
	fmt.Println("Got response: ", string(respBody))
}

func userListUrl() string {
	// Perhaps we should extract the url key into structs to enable Gitea, GitLab and other implementations.
	return cleanUrl(giteaUrlKey, "/api/v1/admin/users")
}
