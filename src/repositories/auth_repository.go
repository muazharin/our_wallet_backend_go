package repositories

import (
	"fmt"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"gorm.io/gorm"
)

type AuthRepo interface {
	CheckAccount(username string, email string, phone string) (int64, database.Users, error)
	SignUp(user database.Users) error
	SignIn(authSignInRequest request.AuthSignInRequest) (database.Users, error)
	SignOut(authSignOutRequest request.AuthSignOutRequest, userId int64) error
	SaveToken(firebaseToken database.FirebaseToken) bool
}

type authConnection struct {
	connection *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authConnection{
		connection: db,
	}
}

func (db *authConnection) CheckAccount(username string, email string, phone string) (int64, database.Users, error) {
	var count int64
	var user database.Users
	db.connection.Where("user_name = ?", username).Or("user_email = ?", email).Or("user_phone = ?", phone).First(&user).Count(&count)

	return count, user, nil
}

func (db *authConnection) SignUp(user database.Users) error {
	res := db.connection.Save(&user)
	if res.Error != nil {
		res.Error = fmt.Errorf("gagal menghubungkan ke database")
		return res.Error
	}
	return nil
}

func (db *authConnection) SignIn(authSignInRequest request.AuthSignInRequest) (database.Users, error) {
	var user database.Users
	res := db.connection.Where("user_name = ?", authSignInRequest.UserName).Or("user_email = ?", authSignInRequest.UserName).Or("user_phone = ?", authSignInRequest.UserName).First(&user)
	if res.Error != nil {
		res.Error = fmt.Errorf("user tidak ditemukan")
		return user, res.Error
	}
	return user, nil
}

func (db *authConnection) SignOut(authSignOutRequest request.AuthSignOutRequest, userId int64) error {
	var firebaseToken database.FirebaseToken
	res := db.connection.Where(&database.FirebaseToken{
		FirebaseTokenUserID: userId,
		FirebaseTokenString: authSignOutRequest.UserFirebaseToken,
	}).Delete(&firebaseToken)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (db *authConnection) SaveToken(firebaseToken database.FirebaseToken) bool {
	res := db.connection.Save(&firebaseToken)
	if res.Error != nil {
		fmt.Println(res.Error)
		return false
	}
	return true
}
