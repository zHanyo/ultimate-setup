// internal/prompt/menu.go
package prompt

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
)

func ShowMenu(options []string) ([]string, error) {
	if len(options) == 0 {
		return nil, errors.New("menu options cannot be empty")
	}

	selected := []string{}
	prompt := &survey.MultiSelect{
		Message: "📦 What do you want to do?",
		Options: options,
	}
	err := survey.AskOne(prompt, &selected)
	if err != nil {
		return nil, err
	}

	if len(selected) == 0 {
		return nil, errors.New("no options selected, please choose at least one")
	}

	return selected, nil
}