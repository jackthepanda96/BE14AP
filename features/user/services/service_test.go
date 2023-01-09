package services

import (
	"api/features/user"
	"api/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	pass string
)

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t)
	inputData := user.Core{Nama: "jerry", Email: "jerr@alterra.id", Alamat: "pasuruan", HP: "08123456", Password: "be1422"}
	resData := user.Core{ID: uint(1), Nama: "jerry", Email: "jerr@alterra.id", Alamat: "pasuruan", HP: "08123456"}
	data.On("Register", mock.Anything).Return(resData, nil).Once()
	srv := New(data)
	res, err := srv.Register(inputData)
	assert.Nil(t, err)
	assert.Equal(t, resData.ID, res.ID)
	assert.Equal(t, resData.Nama, res.Nama)
	data.AssertExpectations(t)
	// srv := New(&mockUserData{})
	// t.Run("Berhasil register", func(t *testing.T) {
	// 	res, err := srv.Register(inputData)
	// 	pass = res.Password
	// 	assert.Nil(t, err)
	// 	assert.Greater(t, res.ID, uint(0))
	// 	// assert.Equal(t, "", res.Password, "")
	// })
}

func TestLogin(t *testing.T) {
	srv := New(&mockUserData{})
	t.Run("Berhasil login", func(t *testing.T) {
		token, res, err := srv.Login("jerry@alterra.id", "be1422")
		assert.Nil(t, err, "err seharusnya nil")
		assert.Greater(t, res.ID, uint(0), "id seharusnya lebih dari 0")
		assert.Equal(t, res.Nama, "jerry")
		assert.NotEmpty(t, token, "seharusnya token tidak kosong")
	})

	t.Run("Password error", func(t *testing.T) {
		srvE := New(&mockUserDataError{})
		token, res, err := srvE.Login("jerry@alterra.id", "be1422")
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)

	})
}

type mockUserData struct{}

func (mud *mockUserData) Login(email string) (user.Core, error) {
	return user.Core{
		ID:       1,
		Nama:     "jerry",
		Email:    "jerry@alterra.id",
		Alamat:   "pasuruan",
		HP:       "0812345678",
		Password: "$2a$10$r8lNLijRdIE3nZpVh60oT.LA.rpl7G9ONQ9RiLCD4fxqZoQLxHIEW",
	}, nil
}
func (mud *mockUserData) Register(newUser user.Core) (user.Core, error) {
	return user.Core{
		ID:     1,
		Nama:   "jerry",
		Email:  "jerry@alterra.id",
		Alamat: "pasuruan",
		HP:     "0812345678",
	}, nil
}
func (mud *mockUserData) Profile(id uint) (user.Core, error) {
	return user.Core{
		ID:     1,
		Nama:   "jerry",
		Email:  "jerry@alterra.id",
		Alamat: "pasuruan",
		HP:     "0812345678",
	}, nil
}

type mockUserDataError struct{}

func (mude *mockUserDataError) Login(email string) (user.Core, error) {
	return user.Core{}, errors.New("terdapat masalah pada server")
}
func (mude *mockUserDataError) Register(newUser user.Core) (user.Core, error) {
	return user.Core{
		ID:     1,
		Nama:   "jerry",
		Email:  "jerry@alterra.id",
		Alamat: "pasuruan",
		HP:     "0812345678",
	}, nil
}
func (mude *mockUserDataError) Profile(id uint) (user.Core, error) {
	return user.Core{
		ID:     1,
		Nama:   "jerry",
		Email:  "jerry@alterra.id",
		Alamat: "pasuruan",
		HP:     "0812345678",
	}, nil
}
