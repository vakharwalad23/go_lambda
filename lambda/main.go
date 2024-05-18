package main

import (
	"fmt"
	"net/http"

	"lambda-function/app"

	"github.com/aws/aws-lambda-go/events"
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
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			return app.ApiHandler.RgisterUserHandler(request)
		case "/login":
			return app.ApiHandler.LoginUser(request)
		default:
			return events.APIGatewayProxyResponse{
				Body:       "Not found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
}
