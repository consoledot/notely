package auth

type (
	Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignIn struct {
		Email           string `json:"email"`
		Name            string `json:"name"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"conform_password"`
	}
)
