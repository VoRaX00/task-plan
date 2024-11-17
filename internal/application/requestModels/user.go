package requestModels

type (
	UserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserRegister struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserToGet struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
