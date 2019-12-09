# ok-cli

# build

```
$ go build -o ok .
```

# add new user

To use this command you need to set the environment variables:

* OK_GITEA_URL -- the url to the Gitea web interface
* OK_GITEA_USERNAME -- the administrative Gitea username
* OK_GITEA_PASSWORD -- the password for `OK_GITEA_USERNAME`
* OK_JENKINS_URL -- the url to the Jenkins web interface
* OK_JENKINS_USERNAME -- the administrative username for Jenkins
* OK_JENKINS_PASSWORD -- the password for `OK_GITLAB_PASSWORD`
* OK_JENKINS_CLI_PATH -- the path to the `jenkins-cli.jar` file (including the jar filename)

# Todo
Those todos need to be put into issues on GitHub. I did not have had the time yet.

* Replace log output with `fmt.Printf` by a appropriate logger from the `log` module.
* Add `ok init <projectName>` to create a repo on Gitea, add the webhooks on Jenkins, create a folder locally with the
  name of `projectName`, initialize a git repo and add a generic Jenkinsfile.
