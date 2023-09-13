package tools

import (
	"github.com/AlecAivazis/survey/v2"
	"xs/pkg/io"
)

func ConfirmDialog(question string) bool {
	var confirm bool
	prompt := &survey.Confirm{
		Message: question,
		Default: true, // Default value (true for "yes")
	}
	err := survey.AskOne(prompt, &confirm, survey.WithValidator(survey.Required))
	if err != nil {
		io.Fatal(err)
	}
	return confirm
}
