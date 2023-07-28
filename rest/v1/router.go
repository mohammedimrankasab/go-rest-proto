package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/mohammedimrankasab/go-rest-proto/handlers"
	"github.com/rs/zerolog/log"
)

const (
	ApiPathEchi     = "/echo"
	ApiPathGetData  = "/get-data"
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
	Charset         = "charset=UTF-8"
	JsonCharset     = ApplicationJson + "; " + Charset
	PingEndpoint    = "/ping"
	PingResponse    = "pong"
)

func loadRoutes(app *handlers.Config) *mux.Router {
	r := mux.NewRouter()

	// Load /ping endpoint
	r.HandleFunc(PingEndpoint, ping)

	r.HandleFunc("/echo", app.Hello).Methods("POST")
	r.HandleFunc("/get-data", app.GetData).Methods("GET", "POST")

	return r
}

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(ContentType, JsonCharset)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(PingResponse); err != nil {
		log.Error().Err(err).Msg("Failed to Write Ping Response")
	}
}
