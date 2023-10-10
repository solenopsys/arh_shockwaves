package publish

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
	"xs/internal/configs"
	jobs_publish "xs/internal/jobs/jobs-publish"
	"xs/pkg/io"
)

var cmdGit = &cobra.Command{
	Use:   "git [repo-url] ",
	Short: "Publish single git repo in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitRepoUrl := args[0]

		split := strings.Split(gitRepoUrl, "/")
		lastSection := split[len(split)-1]
		gitRepoName := strings.Replace(lastSection, ".git", "", 1)
		prefix := strings.Split(gitRepoName, "-")[0]

		group := viper.GetString("git.prefixes." + prefix)
		groupDir := viper.GetString("git.paths." + group)

		nickname := configs.GetAuthManager().Nickname

		ipfsNodeHost := viper.GetString("hosts.ipfsNode")
		pinningServiceHosts := viper.GetString("hosts.pinningService")
		job := jobs_publish.NewPublishGitRepo(ipfsNodeHost, pinningServiceHosts, nickname, group, gitRepoName, groupDir, gitRepoUrl)
		result := job.Execute()
		if result.Error != nil {
			io.Fatal(result.Error)
		} else {
			io.Println(result.Description)
		}
	},
}
