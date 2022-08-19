package repositories

import (
	"fmt"
	"strings"
	"time"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"gorm.io/gorm"
)

type TransRepo interface {
	CreateTransaction(transCreateReq request.TransCreateReq, userId int64) ([]database.TransactionFile, error)
}

type transRepo struct {
	connection *gorm.DB
}

func NewTransRepo(connection *gorm.DB) TransRepo {
	return &transRepo{
		connection: connection,
	}
}

func (db *transRepo) CreateTransaction(transCreateReq request.TransCreateReq, userId int64) ([]database.TransactionFile, error) {
	var transaction database.Transaction
	var trFile database.TransactionFile
	var trFiles []database.TransactionFile
	transaction.TransactionID = time.Now().Unix()
	transaction.TransactionUserID = userId
	transaction.TransactionWalletID = transCreateReq.TransWalletId
	transaction.TransactionType = transCreateReq.TransType
	transaction.TransactionCategory = transCreateReq.TransCategoryId
	transaction.TransactionDetail = transCreateReq.TransDetail
	transaction.TransactionPrice = transCreateReq.TransPrice
	transaction.TransactionDate = time.Now()
	res := db.connection.Save(&transaction)
	if res.Error != nil {
		res.Error = fmt.Errorf("gagal menyimpan transaksi")
		return nil, res.Error
	}

	for i, v := range transCreateReq.TransFile {
		trFile.TfID = time.Now().Unix() + int64(i)
		trFile.TfTransactionID = transaction.TransactionID
		trFile.TfFile = fmt.Sprintf("%v%v.%v", time.Now().Unix(), i, strings.Split(v.Filename, ".")[1])
		trFiles = append(trFiles, trFile)
		res = db.connection.Save(&trFile)
		if res.Error != nil {
			res.Error = fmt.Errorf("gagal menyimpan file transaksi")
			return nil, res.Error
		}
	}
	return trFiles, nil
}
