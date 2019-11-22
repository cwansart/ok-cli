package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
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
		fmt.Printf("An error occurred during request creation: %s\n", err) // TODO: use logger or an error output instead?
		return
	}

	// TODO: handle missing env var
	req.SetBasicAuth(config.Username, config.Password)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("An error occurred during request: %s\n", err)
		return
	}

	respBody, _ := ioutil.ReadAll(resp.Body) // TODO: error handling

	// TODO: pretify JSON output or should we just output the names?
	fmt.Printf("Got response: %s\n", string(respBody))
}

func userListUrl() string {
	// TODO: give an option to get users from Jenkins or Gitea. Or perhaps get them from both?
	// Perhaps we should extract the url key into structs to enable Gitea, GitLab and other implementations.
	return cleanUrl(config.GiteaURL, "/api/v1/admin/users")
}
