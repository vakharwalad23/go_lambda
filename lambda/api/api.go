package api

import (
	"encoding/json"
	"fmt"
	"lambda-function/database"
	"lambda-function/types"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

// Handler to Register new user
func (api ApiHandler) RgisterUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var registerUser types.RgisterUser

	err := json.Unmarshal([]byte(request.Body), &registerUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid Request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	if registerUser.Username == "" || registerUser.Password == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Provide necessary parameters to request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	// First check is user already exists
	userExist, err := api.dbStore.DoesUserExist(registerUser.Username)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal Server error",
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	if userExist {
		return events.APIGatewayProxyResponse{
			Body:       "User already exists with this username",
			StatusCode: http.StatusConflict,
		}, err
	}

	user, err := types.NewUser(registerUser)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal Server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("unable to create new user: %w", err)
	}

	// No existing user logic aka rgister new user
	err = api.dbStore.InsertUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal Server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("error inserting user: %w", err)
	}

	return events.APIGatewayProxyResponse{
		Body:       "Successfully Registered user",
		StatusCode: http.StatusOK,
	}, nil
}
