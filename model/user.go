package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"nama"`
	HP       string `json:"hp" form:"hp"`
	Email    string `json:"email" form:"email"`
	Password string `json:"pwd" form:"password"`
}

type UserModel struct {
	DB *gorm.DB
}

func (um *UserModel) Insert(newUser User) (User, error) {
	err := um.DB.Create(&newUser).Error
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}
