package rbac

type Credentials struct {
	Email    string
	Password string
}

type User struct {
	ID       int64
	Email    string
	Password string
	Role     string
}
