package domain

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(Name string, Email string, Password string) *User {
	return &User{Name: Name, Email: Email, Password: Password}
}
