package publish

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
	"xs/internal/configs"
	"xs/internal/jobs"
	jobs_publish "xs/internal/jobs/jobs-publish"
	"xs/pkg/io"
	"xs/pkg/tools"
)

var cmdSyncGit = &cobra.Command{
	Use:   "sync [config] ",
	Short: "Public file in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configName := args[0]
		nickname := configs.GetAuthManager().Nickname
		conf := configs.GetInstanceConfManager().Conf
		pg := &PublicGit{
			IpfsHost:    conf.Hosts.IpfsHost,
			PinningHost: conf.Hosts.PinningHost,
		}

		err := pg.LoadConfig(configName)

		if err != nil {
			io.Println(err)
			return
		}

		jobsPlan := pg.ManeJobsPlan(nickname)

		for _, job := range jobsPlan {
			jobs.PrintJob((job).Description())
		}

		confirm := tools.ConfirmDialog("Load packets?")

		if confirm {
			io.Println("Proceeding with the action.")
			jobs.ExecuteJobsSync(jobs.ConvertJobs(jobsPlan))

		} else {
			io.Println("Canceled.")
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

func (pg *PublicGit) ManeJobsPlan(nickname string) []jobs.PrintableJob {
	var jobsPlan = make([]jobs.PrintableJob, 0)
	for group, repoNames := range pg.Config.Groups {
		for _, repoName := range repoNames {
			io.Println("Processing repo ", repoName)
			cloneTo := configs.GetInstanceConfManager().Conf.Git.Paths[group]
			repoFullPath := pg.Config.Remote + repoName
			job := jobs_publish.NewPublishGitRepo(pg.IpfsHost, pg.PinningHost, nickname, group, repoName, cloneTo, repoFullPath)
			jobsPlan = append(jobsPlan, job)
		}
	}
	return jobsPlan
}
