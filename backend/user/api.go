package user

import (
	"fmt"
	"net/http"

	cryptolib "github.com/consoledot/notely/utils/crypto"
	"github.com/consoledot/notely/utils/httplib"
)

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	c := httplib.C{W: w, R: r}
	// var tokenResponse auth.TokenResponse
	tokenResponse := r.Context().Value(cryptolib.TokenResponse{}).(cryptolib.TokenResponse)
	fmt.Println(tokenResponse.Id)

	var user User
	result, err := user.GetUser("email", tokenResponse.Email)
	if err != nil {
		c.Response(false, nil, "No user found", http.StatusForbidden, nil)
		return
	}

	c.Response(true, result, "user found", http.StatusOK, nil)

}
