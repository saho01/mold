package parse

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func Location() (string, string, error) {
	fmt.Println("******************************************")

	prompt := promptui.Prompt{
		Label: "Enter project name",
	}
	project, err := prompt.Run()
	if err != nil {
		return "", "", err
	}

	prompt = promptui.Prompt{
		Label: "Enter repo name",
	}
	repo, err := prompt.Run()
	if err != nil {
		return "", "", err
	}
	fmt.Println("******************************************")
	return repo, project, nil

}
