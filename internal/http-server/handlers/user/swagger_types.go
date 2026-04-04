package user

// SignInPayload is the request body for POST /user/sign-in
type SignInPayload struct {
	Login    *string `json:"login"`
	Password *string `json:"password"`
}

// SignUpPayload is the request body for POST /user/sign-up
type SignUpPayload struct {
	Username *string `json:"username"`
	Login    *string `json:"login"`
	Password *string `json:"password"`
}
