package public

import (
	"github.com/spf13/cobra"
	"strings"
	"xs/internal/configs"
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

		conf := configs.GetInstanceConfManager().Conf
		git := conf.Git

		group := git.Prefixes[prefix]
		groupDir := git.Paths[group]

		nickname := configs.GetAuthManager().Nickname

		job := jobs_publish.NewPublishGitRepo(conf.Hosts.IpfsHost, conf.Hosts.PinningHost, nickname, group, gitRepoName, groupDir, gitRepoUrl)
		result := job.Execute()
		if result.Error != nil {
			io.Fatal(result.Error)
		} else {
			io.Println(result.Description)
		}
	},
}
