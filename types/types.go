package types

type DataBase interface {
	CreateUser(user *User) error
}

type User struct {
	UserName string
	Email    string
	Password string
}

func CreateUser(userName string, email string, password string) *User {
	return &User{
		UserName: userName,
		Email:    email,
		Password: password,
	}
}
