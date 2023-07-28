package rest

import (
	"context"
	"net/http"
	"time"

	handlers "github.com/mohammedimrankasab/go-rest-proto/handlers"
	"github.com/rs/zerolog/log"
)

const webServerPort = "8099"

type server struct {
	RestServer      *http.Server
	ShutdownFlag    bool
	SleepTime       time.Duration
	SleepUnit       time.Duration
	RunErrorChannel chan error
}

type Server interface {
	Serve(app *handlers.Config) error
	Stop()
}

func NewServer() Server {
	return &server{
		ShutdownFlag:    false,
		SleepTime:       time.Duration(5),
		SleepUnit:       time.Second,
		RunErrorChannel: make(chan error),
	}
}

func (s *server) startServer() {
	var err error
	// Setup connection to the REST Endpoints
	log.Print("Starting the API server on port ", webServerPort)
	for !s.ShutdownFlag {
		err = s.RestServer.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("failed to start REST server")
			log.Info().Msgf("Waiting for %s...", s.SleepTime*s.SleepUnit)
			time.Sleep(s.SleepUnit * s.SleepTime)
		} else {
			break
		}
	}
}

func (s *server) Serve(app *handlers.Config) error {

	r := loadRoutes(app)

	s.RestServer = &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + webServerPort,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}
	go s.startServer()

	// block forever unless there is a critical error
	return <-s.RunErrorChannel
}

func (s *server) Stop() {
	s.ShutdownFlag = true
	if s.RestServer != nil {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := s.RestServer.Shutdown(ctx)
		if err != nil {
			log.Error().Err(err).Msg("failed to close rest server")
		}
	}
}
