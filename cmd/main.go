package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	handlers "github.com/mohammedimrankasab/go-rest-proto/handlers"
	rest "github.com/mohammedimrankasab/go-rest-proto/rest/v1"
)

func main() {

	app := handlers.Config{}

	gracefulStop := make(chan error, 2)
	listenForInterrupt(gracefulStop)

	var s rest.Server

	go func() {
		s = rest.NewServer()
		gracefulStop <- s.Serve(&app)
	}()

	sig := <-gracefulStop // <-- Blocking call
	log.Info().Msgf("caught Signal: %+v", sig)
	log.Info().Msg("wait for 2 second to finish processing")
	if s != nil {
		s.Stop()
	}
	time.Sleep(2 * time.Second)
}

func listenForInterrupt(errChan chan error) {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
}
