package data

import (
	"api/features/book/data"
	"api/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Email    string
	Alamat   string
	HP       string
	Password string
	Book     []data.Books
}

func ToCore(data User) user.Core {
	return user.Core{
		ID:       data.ID,
		Nama:     data.Nama,
		Email:    data.Email,
		Alamat:   data.Alamat,
		HP:       data.HP,
		Password: data.Password,
	}
}

func CoreToData(data user.Core) User {
	return User{
		Model:    gorm.Model{ID: data.ID},
		Nama:     data.Nama,
		Email:    data.Email,
		Alamat:   data.Alamat,
		HP:       data.HP,
		Password: data.Password,
	}
}
