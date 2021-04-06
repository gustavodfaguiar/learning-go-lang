package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gustavodfaguiar/learning-go-lang/core/task"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "data/task.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := task.NewService(db)
	r := mux.NewRouter()

	n := negroni.New(
		negroni.NewLogger(),
	)

	r.Handle("/v1/task", n.With(
		negroni.Wrap(hello(service)),
	)).Methods("GET", "OPTIONS")
	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":4000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func hello(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tasks, _ := service.GetAll()
		for _, t := range tasks {
			fmt.Println(t)
		}
	})
}
