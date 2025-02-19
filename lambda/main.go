package main

import (
	"fmt"
	"log"
	"net/http"

	"lambda-function/app"
	"lambda-function/middleware"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
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

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func ProtectedHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "This is a protected path",
		StatusCode: http.StatusOK,
	}, nil
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
		case "/protected":
			return middleware.ValidateJWTMiddleware(ProtectedHandler)(request)
		default:
			return events.APIGatewayProxyResponse{
				Body:       "Not found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
}
