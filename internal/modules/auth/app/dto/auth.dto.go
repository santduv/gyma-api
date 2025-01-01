package dto

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	AccessToken string `json:"accessToken"`
}
