package notes

import "github.com/gorilla/mux"

func NotesRouter(router *mux.Router) {
	// Get request
	router.HandleFunc("", CreateNewNotes).Methods("POST")
	router.HandleFunc("", GetNotes).Methods("GET")
	router.HandleFunc("/{id}", DeleteNote).Methods("DELETE")
	router.HandleFunc("/{id}", GetNote).Methods("GET")
	router.HandleFunc("/{id}", EditNote).Methods("PUT")

}
