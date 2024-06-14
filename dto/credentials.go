package dto

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ConfirmUser struct {
	Email            string `json:"email"`
	ConfirmationCode string `json:"confirmationCode"`
}
