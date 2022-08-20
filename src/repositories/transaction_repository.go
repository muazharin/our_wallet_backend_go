package repositories

import (
	"fmt"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"gorm.io/gorm"
)

type TransRepo interface {
	CreateTransaction(transaction database.Transaction) error
	SaveFileTrans(trFile database.TransactionFile) error
	GetFileByTransId(transId int64, page int64) ([]database.TransactionFile, error)
	GetAllTransByWalletId(transWalletId request.TransByWalletIdReq) ([]database.Transaction, error)
	GetAllTransByUserId(transUserId request.TransByUserIdReq) ([]database.Transaction, error)
	GetTransById(transId request.TransByIdReq) (database.Transaction, error)
	CheckIsSeen(transId int64, userId int64) (bool, error)
	SetIsSeen(transIsSeen database.TransactionIsSeen) error
}

type transRepo struct {
	connection *gorm.DB
}

func NewTransRepo(connection *gorm.DB) TransRepo {
	return &transRepo{
		connection: connection,
	}
}

func (db *transRepo) CheckIsSeen(transId int64, userId int64) (bool, error) {
	var transIsSeen database.TransactionIsSeen
	res := db.connection.Where(&database.TransactionIsSeen{
		TransactionID:     transId,
		TransactionUserID: userId,
	}).First(&transIsSeen)

	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}

func (db *transRepo) SetIsSeen(transIsSeen database.TransactionIsSeen) error {
	res := db.connection.Save(&transIsSeen)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (db *transRepo) SaveFileTrans(trFile database.TransactionFile) error {
	res := db.connection.Save(trFile)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (db *transRepo) GetFileByTransId(transId int64, page int64) ([]database.TransactionFile, error) {
	var trFiles []database.TransactionFile
	var res *gorm.DB
	if page != 0 {
		res = db.connection.Model(&database.TransactionFile{}).
			Where("tf_transaction_id = ?", transId).
			Offset((int(page) - 1) * 10).Limit(10).Scan(&trFiles)
	} else {
		res = db.connection.Model(&database.TransactionFile{}).
			Where("tf_transaction_id = ?", transId).
			Scan(&trFiles)
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return trFiles, nil
}

func (db *transRepo) CreateTransaction(transaction database.Transaction) error {
	res := db.connection.Save(&transaction)
	if res.Error != nil {
		res.Error = fmt.Errorf("gagal menyimpan transaksi")
		return res.Error
	}
	return nil
}

func (db *transRepo) GetAllTransByWalletId(transWalletId request.TransByWalletIdReq) ([]database.Transaction, error) {
	var trans []database.Transaction
	res := db.connection.Model(&database.Transaction{}).
		Where("transaction_wallet_id = ?", transWalletId.TransWalletId).
		Order("transaction_id DESC").
		Offset((int(transWalletId.Page) - 1) * 10).Limit(10).Scan(&trans)

	if res.Error != nil {
		return nil, res.Error
	}

	return trans, nil
}

func (db *transRepo) GetAllTransByUserId(transUserId request.TransByUserIdReq) ([]database.Transaction, error) {
	var trans []database.Transaction
	res := db.connection.Model(&database.Transaction{}).
		Where("transaction_user_id = ?", transUserId.TransUserId).
		Order("transaction_id DESC").
		Offset((int(transUserId.Page) - 1) * 10).Limit(10).Scan(&trans)

	if res.Error != nil {
		return nil, res.Error
	}

	return trans, nil
}

func (db *transRepo) GetTransById(transId request.TransByIdReq) (database.Transaction, error) {
	var trans database.Transaction
	res := db.connection.Model(&database.Transaction{}).
		Where("transaction_id = ?", transId.TransId).
		Order("transaction_id DESC").First(&trans)

	if res.Error != nil {
		return database.Transaction{}, res.Error
	}

	return trans, nil
}
