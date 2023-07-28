package pkg

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	hello "github.com/mohammedimrankasab/go-rest-proto/protos"
	"google.golang.org/protobuf/encoding/protojson"
)

type Config struct{}

func (app *Config) Hello(w http.ResponseWriter, r *http.Request) {

	request := &hello.HelloRequest{}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to read message from request ")
	}
	err = protojson.Unmarshal(data, request)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to masrshal from request ")
	}
	message := request.GetMessage()
	result := &hello.HelloResponse{Message: message}
	response, err := protojson.Marshal(result)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to marshal response ")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (app *Config) GetData(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)

	if err != nil {
		log.Panic().Err(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}
