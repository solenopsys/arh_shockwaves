package jobs_publish

import (
	"bytes"
	"errors"
	"os/exec"
	"xs/internal/jobs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

const NPM_APPLICATION = "pnpm"

type PublishFrontLib struct {
	dist string
}

func (t *PublishFrontLib) Execute() *jobs.Result {

	io.Println("Publish library", t.dist)
	pt := xstool.PathTools{}
	pt.SetBasePathPwd()
	pt.MoveTo(t.dist)

	println("CURRENT DIR", pt.GetPwd())

	args := []string{"publish", "--no-git-checks", "--access", "public"}

	cmd := exec.Command(NPM_APPLICATION, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		io.Println(err)
	}

	// Set the Stdout of the command to our buffer

	linkRes := cmd.ProcessState.ExitCode()

	pt.MoveToBasePath()
	if linkRes != 0 {
		o := stdout.String()

		return &jobs.Result{
			Success:     false,
			Err:         errors.New("ERROR PNPM PUBLISH"),
			Description: o,
		}

	}

	return &jobs.Result{
		Success:     true,
		Err:         nil,
		Description: "PublishFrontLib executed",
	}
}
