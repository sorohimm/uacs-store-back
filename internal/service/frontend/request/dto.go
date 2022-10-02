package request

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (o *LoginRequest) SetUsername(username string) *LoginRequest {
	o.Username = username
	return o
}

func (o *LoginRequest) SetPassword(password string) *LoginRequest {
	o.Password = password
	return o
}

func NewRegisterRequest() *RegisterRequest {
	return &RegisterRequest{}
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (o *RegisterRequest) SetUsername(username string) *RegisterRequest {
	o.Username = username
	return o
}

func (o *RegisterRequest) SetPassword(password string) *RegisterRequest {
	o.Password = password
	return o
}

func (o *RegisterRequest) SetEmail(email string) *RegisterRequest {
	o.Email = email
	return o
}

func (o *RegisterRequest) SetRole(role string) *RegisterRequest {
	o.Role = role
	return o
}

type JWT struct {
	AccessToken  string
	RefreshToken string
}
