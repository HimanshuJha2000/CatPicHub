package bootstrap

import (
	"fmt"
	"github.com/DevtronLabs/CatPicHub/internal"
	"github.com/DevtronLabs/CatPicHub/internal/config"
	"github.com/DevtronLabs/CatPicHub/internal/constants"
	"github.com/DevtronLabs/CatPicHub/internal/providers/database"
	"github.com/tylerb/graceful"
	"log"
	"time"
)

// BaseInit Function will be used to load config for both workers and web
func BaseInitAPI(env string) {
	config.LoadConfig(env)
	database.Initialize()

	log.Println("CatPicHub microservice started ...")
	router := internal.SetupRouter()
	err := graceful.RunWithErr(GetListenAddress(), constants.GracefulTimeoutDuration*time.Second, router)
	if err != nil {
		log.Println("Error occurred while starting CatPicHub server ", err)
		panic("Stopping server!!!")
	}
}

// GetListenAddress will give the address in string to listen to
func GetListenAddress() string {
	application := config.GetConfig().Application
	return fmt.Sprintf("%s:%d", application.ListenIP, application.ListenPort)
}
