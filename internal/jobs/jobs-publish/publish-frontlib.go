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

	io.Println("CURRENT DIR", pt.GetPwd())

	args := []string{"publish", "--no-git-checks", "--access", "public"}

	cmd := exec.Command(NPM_APPLICATION, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}

	linkRes := cmd.ProcessState.ExitCode()

	pt.MoveToBasePath()
	if linkRes != 0 {
		o := stdout.String()

		return &jobs.Result{
			Success:     false,
			Error:       errors.New("ERROR PNPM PUBLISH"),
			Description: o,
		}

	}

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "PublishFrontLib executed",
	}
}
