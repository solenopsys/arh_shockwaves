package controllers

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

type Configuration struct {
	Paths  map[string]string   `json:"paths"`
	Groups map[string][]string `json:"groups"`
	Remote string              `json:"remote"`
}

type PublicGit struct {
	IpfsHost    string
	PinningHost string
	config      Configuration
}

func (pg *PublicGit) unpackFiles(gitDir string) {
	const subDile = "objects/pack"

	// scan files and print
	files, err := os.ReadDir(gitDir + "/" + subDile)
	if err != nil {
		io.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".pack") {
			packCurrent := subDile + "/" + file.Name()
			fileBytes, err := os.ReadFile(packCurrent)
			if err != nil {
				io.Fatal(err)
			}
			err = pg.unpackFile(fileBytes)
			if err != nil {
				io.Fatal(err)
			} else {
				err := os.Remove(packCurrent)
				if err != nil {
					io.Fatal(err)
				}
			}
		}

	}
}

func (pg *PublicGit) unpackFile(fileBytes []byte) error {
	cmdName := "git"
	cmdArgs := []string{"unpack-objects", "-q", "-r"}

	cmd := exec.Command(cmdName, cmdArgs...)

	cmdStdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	_, err = cmdStdin.Write(fileBytes)
	if err != nil {
		return err
	}

	cmdStdin.Close()

	return cmd.Wait()
}

func (pg *PublicGit) PublicGitRepo(repoName string, nickname string, moduleType string) error {
	dir := os.TempDir()

	gitTempDir, err := os.MkdirTemp(dir, "xs-git-"+repoName+"-*")
	if err != nil {
		return err
	}

	repoFullPath := pg.config.Remote + repoName

	err = wrappers.CloneGitRepository(repoFullPath, gitTempDir, false, false)
	gitDir := gitTempDir + "/" + ".git"
	err = os.Chdir(gitDir)
	if err != nil {
		return err
	}

	updateServerInfo := exec.Command("git", "update-server-info")
	err = updateServerInfo.Run()
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "rev-parse", "HEAD")

	output, err := cmd.Output()
	if err != nil {
		return err
	}

	commitHash := strings.TrimSpace(string(output))

	//	unpackFiles(gitDir) todo unpack for reuse blocks in ipfs

	cid, err := wrappers.UploadDirToIpfsNode(pg.IpfsHost, gitDir)

	if err != nil {
		return err
	}

	pinning := &wrappers.Pinning{}

	pinning.Host = "http://" + pg.PinningHost
	pinning.UserKey = "alexstorm" // todo remove it

	labels := make(map[string]string)

	namePack := "@" + nickname + "/" + repoName
	labels["source.url"] = pg.config.Remote + repoName
	labels["code.source"] = namePack
	labels["code.type"] = moduleType
	labels["clone.to"] = pg.config.Paths[moduleType]
	labels["commit.hash"] = commitHash

	err = pinning.SmartPin(cid, labels, repoName)

	if err != nil {
		return err
	}
	return nil
}

func (pg *PublicGit) LoadConfig(fileName string) error {
	configFile, err := os.Open(fileName)
	if err != nil {
		io.Println("Error opening config file:", err)
		return err
	}
	defer configFile.Close()

	err = json.NewDecoder(configFile).Decode(&pg.config)
	if err != nil {
		io.Println("Error decoding config:", err)
		return err
	}
	return nil
}

func (pg *PublicGit) ProcessingFile(nickname string) {
	for group, repoNames := range pg.config.Groups {
		for _, repoName := range repoNames {
			io.Println("Processing repo ", repoName)
			err := pg.PublicGitRepo(repoName, nickname, group)
			if err != nil {
				io.Fatal(err)
			}
		}
	}
}
