package app

import (
	"lambda-function/api"
	"lambda-function/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	//Initialize DB store
	//Pass that to apiHandler
	db := database.NewDynamoDBClient()
	apiHandler := api.NewApiHandler(db)

	return App{
		ApiHandler: apiHandler,
	}
}
