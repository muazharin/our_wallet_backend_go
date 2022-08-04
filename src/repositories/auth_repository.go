package repositories

import (
	"fmt"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"gorm.io/gorm"
)

type AuthRepo interface {
	CheckPhone(checkPhoneRequest request.AuthCheckPhoneRequest) (int64, error)
	CheckAccount(username string, email string, phone string) (int64, error)
	SignUp(user database.Users) error
}

type authConnection struct {
	connection *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authConnection{
		connection: db,
	}
}

func (db *authConnection) CheckPhone(checkPhoneRequest request.AuthCheckPhoneRequest) (int64, error) {
	var count int64
	var user database.Users
	db.connection.Where("user_phone = ?", &checkPhoneRequest.Phone).First(&user).Count(&count)

	return count, nil
}

func (db *authConnection) CheckAccount(username string, email string, phone string) (int64, error) {
	var count int64
	var user database.Users
	db.connection.Where("user_name = ?", username).Or("user_email = ?", email).Or("user_phone = ?", phone).First(&user).Count(&count)

	return count, nil
}

func (db *authConnection) SignUp(user database.Users) error {
	res := db.connection.Save(&user)
	if res.Error != nil {
		res.Error = fmt.Errorf("Failed to create a new employee")
		return res.Error
	}
	return nil
}
