package auth

type Credentials struct {
	Username string
	Email    string
	Password string
}

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
	Role     string
}
