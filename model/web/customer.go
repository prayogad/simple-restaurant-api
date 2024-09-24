package web

type CustomerCreateRequest struct {
	Username string `validate:"required,max=100,min=1" json:"username"`
	Password string `validate:"required,max=100,min=1" json:"password"`
}

type CustomerUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Username string `validate:"max=100" json:"username"`
	Password string `validate:"max=100" json:"password"`
}

type CustomerLoginRequest struct {
	Username string `validate:"required,max=100,min=1" json:"username"`
	Password string `validate:"required,max=100,min=1" json:"password"`
}

type CustomerResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
