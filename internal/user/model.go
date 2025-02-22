package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"index"`
	Password string
	Name     string
}

/*func NewUser(email, password, name string) *User {
	user := &User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	return user
}*/
