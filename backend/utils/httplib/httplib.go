package httplib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type C struct {
	W http.ResponseWriter
	R *http.Request
}

type StandardResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Success bool        `json:"success,omitempty"`
	Token   interface{} `json:"token,omitempty"`
}

func responseJSON(res http.ResponseWriter, status int, object interface{}) {
	res.Header().Set("Content-Resource", "application/json")
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	err := json.NewEncoder(res).Encode(object)
	if err != nil {
		return
	}

}

func (c *C) Response(success bool, data interface{}, message string, status int, token interface{}) {

	response := StandardResponse{
		Success: success,
		Message: message,
		Data:    data,
		Token:   token,
	}

	responseJSON(c.W, status, response)

}

func (c *C) GetParamsById(id string) string {
	vars := mux.Vars(c.R)
	return vars[id]
}

func (c *C) GetJSONfromRequestBody(data any) error {
	err := json.NewDecoder(c.R.Body).Decode(data)
	return err
}

func (c *C) GetTokenStringFromHeader() (interface{}, error) {
	authHeader := c.R.Header.Get("Authorization")

	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		token := authHeader[7:]

		return token, nil
	} else {
		return nil, fmt.Errorf("no token found")
	}
}
