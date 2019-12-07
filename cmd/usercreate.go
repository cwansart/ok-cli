package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/spf13/cobra"
)

var userCreateCmd = &cobra.Command{
	Use:   "create --email <email> --username <username> --password <password>",
	Short: "", // TODO: add Short and Long
	Long:  "",
	Run:   userCreate,
}

var (
	email    string
	username string
	password string
)

type createGiteaUserOption struct {
	Email              string `json:"email"`
	Username           string `json:"username"`
	MustChangePassword bool   `json:"must_change_password"`
	Password           string `json:"password"`
}

func init() {
	userCreateCmd.Flags().StringVarP(&email, "email", "e", "", "User's email address")
	userCreateCmd.Flags().StringVarP(&username, "username", "u", "", "User's login name")
	userCreateCmd.Flags().StringVarP(&password, "password", "p", "", "User's password")
	_ = userCreateCmd.MarkFlagRequired("email")
	_ = userCreateCmd.MarkFlagRequired("username")
	_ = userCreateCmd.MarkFlagRequired("password")
}

func userCreate(_ *cobra.Command, _ []string) {
	s := make(chan bool, 2)
	go createOnGitea(s)
	go createOnJenkins(s)

	g, j := <-s
	if !g || !j {
		// TODO: undo changes if one fails
	}
}

func createOnGitea(s chan bool) {
	u := createGiteaUserOption{
		Email:              email,
		Username:           username,
		MustChangePassword: false,
		Password:           password,
	}

	b, err := json.Marshal(u)
	if err != nil {
		_ = fmt.Errorf("An error occurred during marshalling:  %s\n", err)
		return
	}

	req, err := http.NewRequest("POST", giteaUserCreateUrl(), bytes.NewBuffer(b))
	if err != nil {
		_ = fmt.Errorf("An error occurred during request creation: %s\n", err)
		return
	}

	// TODO: perhaps using a token would be better? See /admin​/users​/{username}​/keys route
	req.SetBasicAuth(config.Gitea.Username, config.Gitea.Password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		_ = fmt.Errorf("An error occurred during request: %s\n", err)
		return
	}

	respBody, _ := ioutil.ReadAll(resp.Body) // TODO: error handling

	switch resp.StatusCode {
	case 201:
		fmt.Printf("User has been created. %s\n", string(respBody))
		s <- true
	case 403:
		fmt.Printf("Request failed (403). You are not authorized for this action. %s\n", string(respBody))
		s <- false
	case 422:
		// api validation error; this should not happen except there are changes in the api
		fmt.Printf("Request failed (422). The Gitea api may have changed. %s\n", string(respBody))
	default:
		s <- false
		fmt.Printf("Request failed (%d). %s\n", resp.StatusCode, string(respBody))
	}
}

// To create a new Jenkins we account we need to invoke the jenkins-cli with some parameters. The command looks like:
// echo 'jenkins.model.Jenkins.instance.securityRealm.createAccount("user", "password")' | java -jar $OK_JENKINS_CLI_PATH -s $OK_JENKINS_URL -auth $OK_JENKINS_USERNAME:$OK_JENKINS_PASSWORD groovy =
func createOnJenkins(s chan bool) {
	echo := exec.Command("echo",
		fmt.Sprintf(`'jenkins.model.Jenkins.instance.securityRealm.createAccount("%s", "%s")'`,
			username,
			password))

	cli := exec.Command("java",
		"-jar", config.Jenkins.CliPath,
		"-s", config.Jenkins.Url,
		"-auth", fmt.Sprintf("$%s:$%s", jenkinsUsernameKey, jenkinsPasswordKey),
		// Not sure if the following line works.
		//"-auth", fmt.Sprintf("%s:%s", config.Jenkins.Username, config.Jenkins.Password),
		"help")

	fmt.Printf("Run command: %s | %s", echo.String(), cli.String())

	pipe, err := echo.StdoutPipe()
	defer pipe.Close() // TODO: IntelliJ warns about unhandled exception. How do we handle this?

	if err != nil {
		fmt.Printf("Run jenkins-cli failed. %s\n", err)
		s <- false
		return
	}

	// pipe in the "cli" command into the echo
	cli.Stdin = pipe

	err = echo.Start()
	if err != nil {
		_ = fmt.Errorf("Run echo failed. %s\n", err)
		s <- false
		return
	}

	o, err := cli.Output()

	if err != nil {
		_ = fmt.Errorf("Grabbing output failed. %s\n", err)
		s <- false
		return
	}

	fmt.Printf("Jenkins-cli output: %s", string(o))

	s <- true
}

func giteaUserCreateUrl() string {
	return cleanUrl(config.Gitea.Url, "/api/v1/admin/users")
}
