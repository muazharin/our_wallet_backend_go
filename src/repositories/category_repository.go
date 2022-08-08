package repositories

import (
	"fmt"
	"strconv"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	CheckCategory(title string) (int64, error)
	AddCategory(category database.Category) error
	GetAllCategory(categoryGetAllRequest request.CategoryGetAllRequest) ([]database.Category, error)
}

type categoryRepo struct {
	connection *gorm.DB
}

func NewCategoryRepo(connection *gorm.DB) CategoryRepo {
	return &categoryRepo{
		connection: connection,
	}
}

func (db *categoryRepo) CheckCategory(title string) (int64, error) {
	var count int64
	var category database.Category
	db.connection.Where("category_title = ?", &title).First(&category).Count(&count)
	return count, nil
}

func (db *categoryRepo) AddCategory(category database.Category) error {
	count, _ := db.CheckCategory(category.CategoryTitle)
	if count > 0 {
		err := fmt.Errorf("kategori sudah ada")
		return err
	}
	err := db.connection.Save(&category)
	if err != nil {
		err.Error = fmt.Errorf("gagal menambah kategori")
		return err.Error
	}
	return nil
}

func (db *categoryRepo) GetAllCategory(categoryGetAllRequest request.CategoryGetAllRequest) ([]database.Category, error) {
	var category []database.Category
	userId, _ := strconv.ParseInt(categoryGetAllRequest.CategoryUserId, 10, 64)
	walletId, _ := strconv.ParseInt(categoryGetAllRequest.CategoryWalletId, 10, 64)
	err := db.connection.Where(database.Category{
		CategoryUserID:   userId,
		CategoryWalletID: walletId,
	}).Offset((int(categoryGetAllRequest.CategoryPage) - 1) * 10).Limit(10).Find(&category)

	if err.Error != nil {
		err.Error = fmt.Errorf("gagal menemukan data")
		return nil, err.Error
	}
	return category, nil
}
