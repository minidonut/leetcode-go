package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

//
// Add new problem

func getId(i *Information) {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "What is ID(numeric) of the problem?",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(0)
	}

	i.id = result
}

func getName(i *Information) {
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "What is name of the problem?",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(0)
	}

	i.name = result
}

func confirm() {
	validate := func(input string) error {
		if strings.ToLower(input) == "n" {
			os.Exit(0)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Press key to process",
		Validate:  validate,
		IsConfirm: true,
	}

	_, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(0)
	}
}
