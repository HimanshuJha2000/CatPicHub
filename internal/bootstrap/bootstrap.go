package bootstrap

import (
	"fmt"
	"github.com/DevtronLabs/CatPicHub/internal"
)

// BaseInit Function will be used to load config for both workers and web
func BaseInitAPI() {
	fmt.Println("Starting the players microservice..")
	r := internal.SetupRouter()

	r.Run()
}
