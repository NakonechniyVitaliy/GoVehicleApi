package gearbox

type SignUpDTO struct {
	Username *string `json:"username,omitempty"`
	Login    *string `json:"login,omitempty"`
	Password *string `json:"password,omitempty"`
}

type SignInDTO struct {
	Login    *string `json:"login,omitempty"`
	Password *string `json:"password,omitempty"`
}
