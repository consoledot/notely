package auth

import (
	"context"
	"net/http"

	cryptolib "github.com/consoledot/notely/utils/crypto"
	"github.com/consoledot/notely/utils/httplib"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := httplib.C{W: w, R: r}
		token, err := c.GetTokenStringFromHeader()
	
		if err != nil {
			c.Response(false, nil, "Unauthorized", http.StatusBadRequest, nil)
			return
		}
		userEmail, userId, err := cryptolib.ParseToken((token))
		if err != nil {

			c.Response(false, nil, "Unauthorized", http.StatusUnauthorized, nil)
			return
		}
		emailString, _ := userEmail.(string)
		userIdString, _ := userId.(string)
		tokenResponse := cryptolib.TokenResponse{
			Id:    userIdString,
			Email: emailString,
		}

		ctx := context.WithValue(r.Context(), cryptolib.TokenResponse{}, tokenResponse)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
