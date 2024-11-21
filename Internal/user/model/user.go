package user

type User struct {
	Username string   `json:name`
	Email    string   `json:email`
	Password string   `json:password`
	Roles    []string `json:roles`
}

var Usuarios []User
