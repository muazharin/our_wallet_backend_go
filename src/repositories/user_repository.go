package repositories

import (
	"fmt"
	"log"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo interface {
	CheckUserByID(userId int64) (int64, error)
	CreatePassword(userCreatePasswordRequest request.UserCreatePasswordRequest, userId int64) error
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepo(connection *gorm.DB) UserRepo {
	return &userConnection{
		connection: connection,
	}
}

func (db *userConnection) CheckUserByID(userId int64) (int64, error) {
	var count int64
	var user database.Users
	fmt.Println(userId)
	db.connection.Where("user_id = ?", &userId).First(&user).Count(&count)

	return count, nil
}
func (db *userConnection) CreatePassword(userCreatePasswordRequest request.UserCreatePasswordRequest, userId int64) error {
	var user database.Users
	count, _ := db.CheckUserByID(userId)
	if count < 0 {
		err := fmt.Errorf("user tidak ditemukan")
		return err
	}
	db.connection.Where("user_id = ?", &userId).First(&user)
	user.UserPassword = hashAndSalt([]byte(userCreatePasswordRequest.Password))
	user.UserStatus = "complete"
	err := db.connection.Save(&user)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
