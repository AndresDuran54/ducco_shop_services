package bind

type SessionsLogin struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type SessionsTokenInfo struct {
	Token *string `header:"token"`
}
