package jobs_publish

import (
	"bytes"
	"xs/internal/jobs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

type PublishFrontLib struct {
}

func (t *PublishFrontLib) Execute() *jobs.Result {

	io.Println("Publish library", dist)
	pt := xstool.PathTools{}
	pt.SetBasePathPwd()
	pt.MoveTo(dist)

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

		io.Println("ERROR PNPM PUBLISH: " + o)
		return errors.New("ERROR PNPM PUBLISH")
	}

	return nil
	return &jobs.Result{
		Success:     true,
		Err:         nil,
		Description: "PublishFrontLib executed",
	}
}
