package auth

import "github.com/gorilla/mux"

func AuthRouter(router *mux.Router) {
	router.HandleFunc("/create-account", CreateAccount).Methods("POST")
	router.HandleFunc("/sign-in", SignIn).Methods("POST")
}
