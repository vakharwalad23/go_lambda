package main

import (
	"fmt"

	"lambda-function/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Username string `json:"username"`
}

// This will take payload and do ops on it
func HandleRequest(event Event) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username can not be empty")
	}

	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

func main() {
	//TODO
	app := app.NewApp()
	lambda.Start(app.ApiHandler.RgisterUserHandler)
}
