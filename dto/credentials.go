package dto

type SignUpResendRequest struct {
	Email string `json:"email"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ConfirmUser struct {
	Email            string `json:"email"`
	ConfirmationCode string `json:"confirmationCode"`
}

type LoginResponse struct {
	Status       string `json:"status"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	IdToken      string `json:"idToken"`
	ExpiresIn    int    `json:"expiresIn"`
	CompanyName  string `json:"company_name"`
}
