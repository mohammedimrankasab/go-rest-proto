package pkg

import (
	"io"
	"log"
	"net/http"

	hello "github.com/mohammedimrankasab/go-rest-proto/protos"
	"google.golang.org/protobuf/encoding/protojson"
)

type Config struct {
}

func (app *Config) Hello(w http.ResponseWriter, r *http.Request) {

	request := &hello.HelloRequest{}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}
	err = protojson.Unmarshal(data, request)
	if err != nil {
		log.Fatalf("Unable to masrshal from request : %v", err)
	}
	message := request.GetMessage()
	result := &hello.HelloResponse{Message: message}
	response, err := protojson.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
