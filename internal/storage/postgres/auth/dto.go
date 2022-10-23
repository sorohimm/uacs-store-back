// Package auth TODO
package auth

type Credentials struct {
	UserID   int64
	Username string
	Email    string
	Password string
	PwdSalt  string
}

func NewUser() *User {
	return &User{}
}

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
	Role     string
}

func (o *User) SetID(id int64) *User {
	o.ID = id
	return o
}

func (o *User) SetUsername(un string) *User {
	o.Username = un
	return o
}

func (o *User) SetEmail(email string) *User {
	o.Email = email
	return o
}

func (o *User) SetPassword(pwd string) *User {
	o.Password = pwd
	return o
}

func (o *User) SetRole(r string) *User {
	o.Role = r
	return o
}

type LoginRequest struct {
	ID       int64
	Username string
	Password string
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{}
}

type CreateUserRequest struct {
	User    User
	PwdSalt string
}

func (o *CreateUserRequest) SetUser(user User) *CreateUserRequest {
	o.User = user
	return o
}

func (o *CreateUserRequest) SetSalt(salt string) *CreateUserRequest {
	o.PwdSalt = salt
	return o
}
