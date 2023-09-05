package public

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

var cmdGit = &cobra.Command{
	Use:   "git [path] ",
	Short: "Public file in ipfs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		remoteRepo := args[0]

		// git clone to temp dir

		dir := os.TempDir()

		gitTempDir, err := os.MkdirTemp(dir, "xs-git-*")
		if err != nil {
			log.Fatalf("failed to create temp file: %v", err)
		}

		println(remoteRepo)
		println(gitTempDir)

		err = wrappers.CloneGitRepository(remoteRepo, gitTempDir, false, false)
		gitDir := gitTempDir + "/" + ".git"
		err = os.Chdir(gitDir)
		if err != nil {
			log.Fatal(err)
		}

		cmd1 := exec.Command("git", "update-server-info")
		err = cmd1.Run()
		if err != nil {
			log.Fatal(err)
		}

		//	unpackFiles(gitDir)

		cid, err := wrappers.UploadDirToIpfsNode(IpfsHost, gitDir)

		if err != nil {
			io.Println(err)
		} else {
			io.Println("File cid: ", cid)
		}

		pinning := &wrappers.Pinning{}

		pinning.Host = "http://" + PinningHost

		remoteRepoSplit := strings.Split(remoteRepo, "github.com/") // todo remove
		namePack := "@" + remoteRepoSplit[1]

		labels := make(map[string]string)
		labels["source.url"] = remoteRepo

		labels["code.front-module"] = namePack
		pin, err := pinning.SimplePin(cid, labels)
		println("pin", pin)

		ipns, err := pinning.SetName(cid, namePack)
		if err != nil {
			log.Fatal(err)
		}

		println("ipns", ipns)
		// add to repo ipfs
		// rm temp dir
		// public ipns or update
	},
}

func unpackFiles(gitDir string) {
	const subDile = "objects/pack"

	// scan files and print
	files, err := os.ReadDir(gitDir + "/" + subDile)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// end width .pack
		if strings.HasSuffix(file.Name(), ".pack") {
			println(file.Name())
			packCurrent := subDile + "/" + file.Name()
			fileBytes, err := os.ReadFile(packCurrent)
			if err != nil {
				log.Fatal(err)
			}
			err = unpackFile(fileBytes)
			if err != nil {
				log.Fatal(err)
			} else {
				//err := os.Remove(packCurrent)
				//if err != nil {
				//	log.Fatal(err)
				//}
			}
		}

	}
}

func unpackFile(fileBytes []byte) error {
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
