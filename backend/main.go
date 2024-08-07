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

	"github.com/consoledot/notely/auth"
	db "github.com/consoledot/notely/database"
	"github.com/consoledot/notely/notes"
	"github.com/consoledot/notely/user"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db.Connect()
	defer db.Disconnect()

	server := &http.Server{
		Addr:         ":8181",
		Handler:      routes(),
		ReadTimeout:  50 * time.Second,
		WriteTimeout: 50 * time.Second,
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
	userRouter := router.PathPrefix("/user").Subrouter()
	notesRouter := router.PathPrefix("/notes").Subrouter()
	authRouter := router.PathPrefix("/auth").Subrouter()

	protectedRoutes := []*mux.Router{notesRouter, userRouter}

	for _, routes := range protectedRoutes {
		routes.Use(auth.AuthMiddleware)
	}

	notes.NotesRouter(notesRouter)
	user.UserRoutes(userRouter)
	auth.AuthRouter(authRouter)

	return router
}

func shutdownServer(server *http.Server) {
	// Listening for termination signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("\n Shutting down the server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	// Shutdown server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown", err)
	}

}
