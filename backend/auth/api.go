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
	var request SignUp
	if err := c.GetJSONfromRequestBody(&request); err != nil {
		fmt.Println(err)
		c.Response(false, nil, "Error reading request", http.StatusBadRequest, nil)
		return
	}

	if request.Password != request.ConfirmPassword {
		c.Response(false, nil, "Password must be the same", http.StatusBadRequest, nil)
		return
	}
	passwordHash, err := cryptolib.Hash(request.Password)
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "Error creating account, try again", http.StatusBadRequest, nil)
		return
	}
	user := user.NewUser(request.Email, request.Name, passwordHash)

	if ok := user.DoesUserExit(); ok {
		c.Response(false, nil, "User already exist", http.StatusBadRequest, nil)
		return
	}
	userId, err := user.CreateUser()
	if err != nil {
		fmt.Println(err)
		c.Response(false, nil, "Error creating account", http.StatusBadGateway, nil)
		return
	}

	fmt.Printf("Id: %v", userId)
	token, err := cryptolib.CreateToken(user.Email, userId)
	if err != nil {
		c.Response(false, nil, "Error creating user token, try login in", http.StatusBadGateway, nil)
		return
	}

	c.Response(true, nil, "Request successful", http.StatusCreated, token)

}

func SignIn(w http.ResponseWriter, r *http.Request) {
	c := httplib.C{W: w, R: r}
	var request Login

	if err := c.GetJSONfromRequestBody(&request); err != nil {
		fmt.Println(err)
		c.Response(false, nil, "Error reading request", http.StatusBadRequest, nil)
		return
	}

	user := user.NewUser(request.Email, "", request.Password)

	if ok := user.DoesUserExit(); !ok {
		c.Response(false, nil, "User does not exist", http.StatusBadRequest, nil)
		return

	}

	if ok := user.DoesPassWordMatch(); !ok {
		c.Response(false, nil, "Invalid credentials", http.StatusBadRequest, nil)
		return
	}
	userId := user.Id.String()
	fmt.Printf("userId %v", userId)
	token, err := cryptolib.CreateToken(user.Email, userId)
	if err != nil {
		fmt.Println(err, token)
		c.Response(false, nil, "Error creating user token, try login in", http.StatusBadGateway, nil)
		return
	}
	c.Response(true, nil, "user account gotten", http.StatusOK, token)

}
