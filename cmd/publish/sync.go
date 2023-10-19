package publish

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"xs/internal/configs"
	"xs/internal/jobs"
	jobs_publish "xs/internal/jobs/jobs-publish"
	"xs/pkg/io"
	"xs/pkg/ui"
)

var cmdSyncGit = &cobra.Command{
	Use:   "sync [config]",
	Short: "Mass publish repos from git in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configName := args[0]
		filter := ""
		if len(args) > 1 {
			filter = args[1]
		}

		nickname := configs.GetAuthManager().Nickname
		pg := &PublicGit{
			IpfsHost:    viper.GetString("hosts.ipfsNode"),
			PinningHost: viper.GetString("hosts.pinningService"),
		}

		err := pg.LoadConfig(configName)

		if err != nil {
			io.Println(err)
			return
		}

		jobsPlan := jobs.ConvertPjToJi(pg.ManeJobsPlan(nickname, "*"))
		applied, changedFilter := ui.FilteredListView(jobsPlan, "Publishing git repos", filter)
		if applied {
			jobsPlanApplied := pg.ManeJobsPlan(nickname, changedFilter)
			executor := jobs.NewExecutor(jobsPlanApplied)
			ui.ProcessingJobs(executor)
		}
	},
}

type Configuration struct {
	Groups map[string][]string `yaml:"groups"`
	Remote string              `yaml:"remote"`
}

type PublicGit struct {
	IpfsHost    string
	PinningHost string
	Config      Configuration
}

func (pg *PublicGit) LoadConfig(fileName string) error {
	configFile, err := os.Open(fileName)
	if err != nil {
		io.Println("Error opening config file:", err)
		return err
	}
	defer configFile.Close()

	err = yaml.NewDecoder(configFile).Decode(&pg.Config)
	if err != nil {
		io.Println("Error decoding config:", err)
		return err
	}
	return nil
}

func (pg *PublicGit) ManeJobsPlan(nickname string, filter string) []jobs.PrintableJob {
	var jobsPlan = make([]jobs.PrintableJob, 0)
	for group, repoNames := range pg.Config.Groups {
		for _, repoName := range repoNames {

			if filter != "" {
				matched, err := configs.PatternMatching(repoName, filter)
				if err != nil {
					io.Println("Error:", err)
					continue
				}
				if !matched {
					continue
				}
			}
			cloneTo := viper.GetString("git.paths." + group)
			repoFullPath := pg.Config.Remote + repoName
			job := jobs_publish.NewPublishGitRepo(pg.IpfsHost, pg.PinningHost, nickname, group, repoName, cloneTo, repoFullPath)
			jobsPlan = append(jobsPlan, job)
		}
	}

	return jobsPlan
}
