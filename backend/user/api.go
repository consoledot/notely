package user

import (
	"net/http"

	"github.com/consoledot/notely/utils/httplib"
)

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	c := httplib.C{W: w, R: r}
	userEmail, ok := r.Context().Value("userEmail").(string)

	if !ok {
		c.Response(false, nil, "Invalid user email format", http.StatusInternalServerError, nil)
		return
	}

	var user User
	result, err := user.GetUser("email", userEmail)
	if err != nil {
		c.Response(false, nil, "No user found", http.StatusForbidden, nil)
		return
	}

	c.Response(true, result, "user found", http.StatusOK, nil)

}
