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
		gitRepoName := strings.Replace(".git", split[len(split)-1], "", 1)
		prefix := strings.Split(gitRepoUrl, "_")[0]
		group := PREFIXES[prefix]
		cloneTo := PATHS[group]
		nickname := "solenopsys" // todo get from login

		job := jobs_publish.NewPublishGitRepo(IpfsHost, PinningHost, nickname, group, gitRepoName, cloneTo)
		result := job.Execute()
		if result.Error != nil {
			io.Fatal(result.Error)
		} else {
			io.Println(result.Description)
		}
	},
}
