package auth

type (
	Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignUp struct {
		Email           string `json:"email"`
		Name            string `json:"name"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
)
