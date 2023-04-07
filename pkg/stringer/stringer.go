package stringer

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

var qs = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "What is the name of the project ?"},
		Validate:  survey.Required,
		Transform: survey.ToLower,
	},
	{
		Name: "framework",
		Prompt: &survey.Select{
			Message: "Choose a framework",
			Options: []string{"NextJS", "NestJS"},
			Default: "NextJS",
		},
	},
	{
		Name: "utils",
		Prompt: &survey.MultiSelect{
			Message: "Choose the required utils for the project",
			Options: []string{"Tests", "Auth", "Database Integration", "Make me a coffee 🫖"},
		},
	},
}

func HandleRootCommand() {
	fmt.Println("Running command")
	answers := struct {
		Name      string
		Framework string
		Utils     []string
	}{}

	err := survey.Ask(qs, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", answers)
}

func HandleVersionCommand() {
	fmt.Println("The current version is 0.42")
}
