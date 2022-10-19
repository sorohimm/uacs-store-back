// Package auth TODO
package auth

type Credentials struct {
	UserID   int64
	Username string
	Email    string
	Password string
	PwdSalt  string
}

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
	Role     string
}

type LoginRequest struct {
	ID       int64
	Username string
	Password string
}

type CreateUserRequest struct {
	User    User
	PwdSalt string
}
