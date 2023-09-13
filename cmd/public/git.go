package public

import (
	"github.com/spf13/cobra"
	"strings"
	jobs_publish "xs/internal/jobs/jobs-publish"
	"xs/pkg/io"
)

var cmdGit = &cobra.Command{
	Use:   "git [repo-url] ",
	Short: "Public file in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitRepoUrl := args[0]

		split := strings.Split(gitRepoUrl, "/")
		lastSection := split[len(split)-1]
		gitRepoName := strings.Replace(lastSection, ".git", "", 1)
		prefix := strings.Split(gitRepoName, "-")[0]
		group := PREFIXES[prefix]
		groupDir := PATHS[group]

		nickname := "solenopsys" // todo get from login
		cloneTo := groupDir + "/" + nickname + "/" + gitRepoName

		job := jobs_publish.NewPublishGitRepo(IpfsHost, PinningHost, nickname, group, gitRepoName, cloneTo, gitRepoUrl)
		result := job.Execute()
		if result.Error != nil {
			io.Fatal(result.Error)
		} else {
			io.Println(result.Description)
		}
	},
}
