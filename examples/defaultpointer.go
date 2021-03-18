package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func main() {

	var answers = &struct {
		Input  string
		Output string
	}{}

	defaultAnswer := "something"

	q1 := &survey.Question{
		Name: "input",
		Prompt: &survey.Input{
			Message: "Input?",
			Default: defaultAnswer,
		},
		Validate: func(val interface{}) error {
			input := val.(string)
			defaultAnswer = input
			return nil
		},
	}

	q2 := &survey.Question{
		Name: "output",
		Prompt: &survey.Input{
			Message: "Output?",
			DefaultPointer: &defaultAnswer,
		},
		Validate: survey.Required,
	}

	// the questions to ask
	var questions = []*survey.Question{q1, q2}

	// ask the question
	err := survey.Ask(questions, answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("%+v\n", answers)
}
