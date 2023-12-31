package main

import (
	"flag"
	"github.com/DevtronLabs/CatPicHub/internal/bootstrap"
	"github.com/DevtronLabs/CatPicHub/internal/constants"
	"os"
	"os/signal"
	"syscall"
)

var env *string

// fetch all the cli inputs options provided.
func init() {
	env = flag.String(constants.Env, constants.Development, "Application env : prod/dev")
}

func main() {
	flag.Parse()
	done := make(chan struct{})

	// Wait for termination signal (SIGINT, SIGTERM, SIGKILL)
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-signalChannel
		close(done)
	}()

	bootstrap.BaseInitAPI(*env)

	<-done
}
