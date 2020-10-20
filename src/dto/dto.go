package dto

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CredentialsResponse struct {
	Prefix string `json:"prefix"`
	Token  string `json:"token"`
}
