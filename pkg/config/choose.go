package config

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
	"github.com/xero-github/xoauth/pkg/db"
	"log"
)



func ChooseClient() (string, error) {
	allClients, err := db.GetClients()

	if err != nil {
		panic(err)
	}

	if len(allClients) == 0 {
		log.Fatalf("Please create a connection using `%s`",
			color.Green.Sprintf("xoauth setup [connectionName]"))
	}

	var connections []string

	for _, value := range allClients {
		connections = append(connections, value.Alias)
	}

	connectionPicker := &survey.Select{
		Message: "Choose a client",
		Options: connections,
	}

	var chosenConnection string

	askErr := survey.AskOne(connectionPicker, &chosenConnection, survey.WithValidator(survey.Required))

	if askErr != nil {
		return "", askErr
	}

	return chosenConnection, nil
}

