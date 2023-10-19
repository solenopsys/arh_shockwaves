package jobs_publish

import (
	"os"
	"os/exec"
	"strings"
	"xs/internal/configs"
	"xs/internal/jobs"
	"xs/pkg/wrappers"
	"xs/pkg/wrappers/git"
)

type PublishGitRepo struct {
	nickname    string
	moduleType  string
	repoName    string
	cloneTo     string
	ipfsNode    *wrappers.IpfsNode
	pinningHost string
	steps       []func() error
	dist        string
	gitTempDir  string
	gitDir      string
	commitHash  string
	cid         string
	repoUrl     string
}

func NewPublishGitRepo(ipfsHost string, pinningHost string, nickname string, moduleType string, repoName string, cloneTo string, repoUrl string) *PublishGitRepo {

	return &PublishGitRepo{
		repoUrl:     repoUrl,
		nickname:    nickname,
		moduleType:  moduleType,
		repoName:    repoName,
		cloneTo:     cloneTo,
		ipfsNode:    &wrappers.IpfsNode{ipfsHost},
		pinningHost: pinningHost}
}

func (t *PublishGitRepo) makeTempDir() error {
	var err error
	dir := os.TempDir()
	t.gitTempDir, err = os.MkdirTemp(dir, "xs-git-"+t.repoName+"-*")
	return err
}

func (t *PublishGitRepo) cloneRepository() error {
	err := git.CloneGitRepository(t.repoUrl, t.gitTempDir, true, false, false, "")
	t.gitDir = t.gitTempDir //+ "/" + ".git"
	return err
}

func (t *PublishGitRepo) changeCurrentDir() error {
	return os.Chdir(t.gitDir)
}

func (t *PublishGitRepo) updateServerInfo() error {
	updateServerInfo := exec.Command("git", "update-server-info")
	return updateServerInfo.Run()
}

func (t *PublishGitRepo) extractCommitHash() error {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	output, err := cmd.Output()
	t.commitHash = strings.TrimSpace(string(output))
	return err
}

func (t *PublishGitRepo) publishDirInIpfs() error {
	var err error
	t.cid, err = t.ipfsNode.UploadDirToIpfsNode(t.gitDir)
	return err
}

func (t *PublishGitRepo) unpackFiles() error {
	const subDile = "objects/pack"

	// scan files and print
	files, err := os.ReadDir(t.gitDir + "/" + subDile)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".pack") {
			packCurrent := subDile + "/" + file.Name()
			fileBytes, err := os.ReadFile(packCurrent)
			if err != nil {
				return err
			}
			err = t.unpackFile(fileBytes)
			if err != nil {
				return err
			} else {
				err := os.Remove(packCurrent)
				if err != nil {
					return err
				}
			}
		}

	}
	return nil
}

func (t *PublishGitRepo) unpackFile(fileBytes []byte) error {
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

func (t *PublishGitRepo) pinCidInPinningService() error {
	pinning := &wrappers.Pinning{}

	pinning.Host = t.pinningHost
	pinning.UserKey = configs.GetAuthManager().PublicKey

	labels := make(map[string]string)

	//todo  move to method of  pinning request
	namePack := "@" + t.nickname + "/" + t.repoName
	labels["source.url"] = t.repoUrl
	labels["code.source"] = namePack
	labels["code.type"] = t.moduleType
	labels["clone.to"] = t.cloneTo
	labels["commit.hash"] = t.commitHash
	return pinning.SmartPinAndName(t.cid, labels, t.repoName)
}

func (t *PublishGitRepo) Execute() *jobs.Result {
	t.steps = []func() error{
		t.makeTempDir,
		t.cloneRepository,
		t.changeCurrentDir,
		t.updateServerInfo,
		t.extractCommitHash,
		//	unpackFiles  // todo unpack for reuse blocks in ipfs
		t.publishDirInIpfs,
		t.pinCidInPinningService,
	}

	for _, step := range t.steps {
		err := step()
		if err != nil {
			return &jobs.Result{
				Success: false,
				Error:   err,
			}
		}
	}

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "Git repository published: " + t.repoName,
	}
}

func (t *PublishGitRepo) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: t.repoUrl,
		Name:        t.repoName,
		Key:         "publish-git-repo-" + t.repoName,
	}
}
