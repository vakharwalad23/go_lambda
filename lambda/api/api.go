package api

import (
	"fmt"
	"lambda-function/database"
	"lambda-function/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

// Handler to Register new user
func (api ApiHandler) RgisterUserHandler(event types.RgisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("provide necessary parameters to request")
	}

	// First check is user already exists
	userExist, err := api.dbStore.DoesUserExist(event.Username)

	if err != nil {
		return fmt.Errorf("something went wrong: %w", err)
	}

	if userExist {
		return fmt.Errorf("user already exists with this username")
	}

	// No existing user logic aka rgister new user
	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("error registering the user: %w", err)
	}

	return nil
}
