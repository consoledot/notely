package user

import (
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {

	router.HandleFunc("/me", GetUserDetails).Methods("GET")

}
