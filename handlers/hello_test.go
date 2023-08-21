package handlers

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"

	hello "github.com/mohammedimrankasab/go-rest-proto/protos"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestHello(t *testing.T) {
	recorder := httptest.NewRecorder()
	jsonBody := []byte(`{"message": "hello!"}`)

	app := &Config{}
	request := httptest.NewRequest("POST", "/echo", bytes.NewReader(jsonBody))

	app.Hello(recorder, request)
	res := recorder.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	response := &hello.HelloResponse{}
	err = protojson.Unmarshal(data, response)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to masrshal from response ")
	}

	if response.Message != "hello!" {
		t.Errorf("expected hello! got %v", response.Message)
	}

}

func TestGetDataGet(t *testing.T) {
	recorder := httptest.NewRecorder()
	jsonBody := []byte(`{"message": "hello world!"}`)

	app := &Config{}
	request := httptest.NewRequest("GET", "/get-data", bytes.NewReader(jsonBody))
	app.GetData(recorder, request)
	res := recorder.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if !bytes.Equal(jsonBody, data) {
		t.Errorf("expected hello world! got %v", string(data))
	}
}

func TestGetDataPost(t *testing.T) {
	recorder := httptest.NewRecorder()
	jsonBody := []byte(`{"message": "hello world!"}`)

	app := &Config{}
	request := httptest.NewRequest("POST", "/get-data", bytes.NewReader(jsonBody))
	app.GetData(recorder, request)
	res := recorder.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if !bytes.Equal(jsonBody, data) {
		t.Errorf("expected hello world! got %v", string(data))
	}
}
