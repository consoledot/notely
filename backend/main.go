package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	db "github.com/consoledot/notely/database"
	"github.com/consoledot/notely/notes"
	"github.com/gorilla/mux"
)

func main() {

	db.Connect()
	defer db.Disconnect()

	server := &http.Server{
		Addr:         ":8181",
		Handler:      routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// Graceful shutdown
	go shutdownServer(server)

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

func shutdownServer(server *http.Server) {
	// Listening for termination signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("\n Shutting down the server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// Shutdown server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown", err)
	}

}
