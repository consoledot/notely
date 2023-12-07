package main

import (
	"net/http"

	"github.com/consoledot/notely/notes"
)




func main(){
		server := http.NewServeMux()

		server.HandleFunc("/", notes.NotesController)

		err := http.ListenAndServe(":8080", server)
		if err != nil{
			panic("Server failed to start")
		}
}