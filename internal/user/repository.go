package user

import "server/pkg/db"

type UserRepository struct {
	DataBase *db.Db
}

func NewUserRepository(dataBase *db.Db) *UserRepository {
	return &UserRepository{
		DataBase: dataBase,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.DataBase.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.DataBase.DB.First(&user, "email=?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
