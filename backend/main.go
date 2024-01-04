package main

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/consoledot/notely/database"
	"github.com/consoledot/notely/notes"
	"github.com/gorilla/mux"
)

func main() {

	db.Connect()
	defer db.Disconnect()
	server := &http.Server{
		Addr:    ":8181",
		Handler: routes(),
	}

	fmt.Println("Server listening on port 8181")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln("Server Crashed ", err)
	}

}

func routes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", notes.CreateNewNotes).Methods("POST")
	router.HandleFunc("/", notes.GetNotes).Methods("GET")
	return router
}
