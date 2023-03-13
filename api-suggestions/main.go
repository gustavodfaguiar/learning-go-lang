package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgorilla/v2"
	"go.elastic.co/apm/v2"
)

type ResponseSuggestion struct {
	Suggestions []suggestion `json:"suggestions"`
}

type suggestion struct {
	Text string `json:"text"`
}

type PageData struct {
	TraceID string
	SpanID  string
}

func HandleSuggestions(w http.ResponseWriter, r *http.Request) {
	log.Info("HandleSuggestions")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data := ResponseSuggestion{}
	suggestions := make([]suggestion, 0)

	suggestions = append(suggestions, suggestion{
		Text: "Hotel Caiuá",
	})

	data.Suggestions = suggestions

	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

	return
}

func main() {
	os.Setenv("ELASTIC_APM_SERVICE_NAME", "search-suggestion")
	os.Setenv("ELASTIC_APM_SERVER_URL", "http://localhost:8200")

	r := mux.NewRouter()
	r.HandleFunc("/suggestions", HandleSuggestions)
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/layout.html"))
		tmpl.Execute(w, apm.TransactionFromContext(r.Context()))
	})
	apmgorilla.Instrument(r)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
