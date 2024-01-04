package main

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/consoledot/notely/database"
	"github.com/consoledot/notely/notes"
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
	router := http.NewServeMux()

	router.HandleFunc("/", notes.NotesController)
	return router
}
