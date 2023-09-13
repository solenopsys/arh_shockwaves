package jobs_env

import (
	"os/exec"
	"strings"
	"xs/internal/jobs"
	"xs/pkg/io"
)

type AppCheck struct {
	command string
	params  []string
}

func (t *AppCheck) Execute() *jobs.Result {
	arg := t.params[1]
	splitArg := strings.Split(arg, " ")
	var version, err = exec.Command(t.params[0], splitArg...).Output()
	if err == nil {
		verLine := string(version)
		verLine = strings.Replace(verLine, "version", "", 1)
		verLine = strings.Replace(verLine, t.command, "", 1)
		verLine = strings.TrimSpace(verLine)
		io.Println("")
		io.Println(t.command)
		io.Println(" -------------------------------->")
		io.Println(verLine)
		return &jobs.Result{
			Success:     true,
			Error:       nil,
			Description: "AppCheck executed",
		}
	} else {
		return &jobs.Result{
			Success:     true,
			Error:       err,
			Description: t.command + ":not installed",
		}
	}
}

func NewAppCheck(command string, params []string) *AppCheck {
	return &AppCheck{command: command, params: params}
}
