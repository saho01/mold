package parse

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func Location() (string, error) {
	fmt.Println("******************************************")

	prompt := promptui.Prompt{
		Label: "Enter project name",
	}
	project, err := prompt.Run()
	if err != nil {
		return "", err
	}
	fmt.Printf("porject name :%s\n", project)

	prompt = promptui.Prompt{
		Label: "Enter dir",
	}
	dir, err := prompt.Run()
	if err != nil {
		return "", err
	}
	fmt.Println("******************************************")
	return fmt.Sprintf("%s/%s", dir, project), nil

}
