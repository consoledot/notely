package auth

import (
	"fmt"
	"net/http"

	"github.com/consoledot/notely/user"
	cryptolib "github.com/consoledot/notely/utils/crypto"
	"github.com/consoledot/notely/utils/httplib"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	c := httplib.C{W: w, R: r}
	var request SignIn
	if err := c.GetJSONfromRequestBody(&request); err != nil {
		fmt.Println(err)
		c.Response(false, nil, "Error reading request", http.StatusBadRequest)
		return
	}

	if request.Password != request.ConfirmPassword {
		c.Response(false, nil, "Password must be the same", http.StatusBadRequest)
		return
	}
	passwordHash, err := cryptolib.Hash(request.Password)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "Error creating account, try again", http.StatusBadRequest)
		return
	}
	user := user.User{Email: request.Email, Name: request.Name, PasswordHash: passwordHash}

	if ok := user.DoesUserExit(); !ok {
		c.Response(false, nil, "User already exist", http.StatusBadRequest)
		return
	}
	if err := user.CreateUser(); err != nil {
		fmt.Println(err)
		c.Response(false, nil, "Error creating account", http.StatusBadGateway)
		return
	}
	c.Response(true, request, "Request successful", http.StatusCreated)

}
