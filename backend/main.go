package main

import (
	"net/http"
)




func main(){
		server := http.NewServeMux()

		server.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
				w.Write([]byte("Welcome"))
		})

		err := http.ListenAndServe(":8080", server)
		if err == nil{
			panic("Server failed to start")
		}
}