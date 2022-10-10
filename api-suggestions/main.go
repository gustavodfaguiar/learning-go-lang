package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ResponseSuggestion struct {
	Suggestions []suggestion `json:"suggestions"`
}

type suggestion struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/suggestions", func(w http.ResponseWriter, r *http.Request) {
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
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)

		return
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
