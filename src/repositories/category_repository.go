package repositories

import (
	"fmt"
	"strconv"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	CheckCategory(title string, walletId int64) (int64, database.Category, error)
	AddCategory(category database.Category) error
	UpdateCategory(category database.Category) error
	DeleteCategory(category database.Category) error
	GetAllCategory(categoryGetAllRequest request.CategoryGetAllRequest) ([]database.Category, error)
	GetCategoryById(categoryId int64) (database.Category, error)
}

type categoryRepo struct {
	connection *gorm.DB
}

func NewCategoryRepo(connection *gorm.DB) CategoryRepo {
	return &categoryRepo{
		connection: connection,
	}
}

func (db *categoryRepo) CheckCategory(title string, walletId int64) (int64, database.Category, error) {
	var count int64
	var category database.Category
	db.connection.Where(&database.Category{
		CategoryTitle:    title,
		CategoryWalletID: walletId,
	}).First(&category).Count(&count)
	return count, category, nil
}

func (db *categoryRepo) GetCategoryById(categoryId int64) (database.Category, error) {
	var category database.Category
	res := db.connection.Model(&database.Category{}).
		Where("category_id = ?", categoryId).
		First(&category)

	if res.Error != nil {
		return database.Category{}, res.Error
	}

	return category, nil
}

func (db *categoryRepo) AddCategory(category database.Category) error {
	count, res, _ := db.CheckCategory(category.CategoryTitle, category.CategoryWalletID)
	if count > 0 {
		if !res.CategoryIsActive {
			res.CategoryIsActive = true
			err := db.connection.Save(res)
			if err.Error != nil {
				return err.Error
			}
			return nil
		}
		err := fmt.Errorf("kategori sudah ada")
		return err
	}
	err := db.connection.Save(&category)
	if err.Error != nil {
		err.Error = fmt.Errorf("gagal menambah kategori")
		return err.Error
	}
	return nil
}

func (db *categoryRepo) UpdateCategory(category database.Category) error {
	count, _, _ := db.CheckCategory(category.CategoryTitle, category.CategoryWalletID)
	if count > 0 {
		err := fmt.Errorf("kategori sudah ada")
		return err
	}

	res := db.connection.Save(&category)
	if res.Error != nil {
		res.Error = fmt.Errorf("gagal mengupdate kategori")
		return res.Error
	}
	return nil
}

func (db *categoryRepo) DeleteCategory(category database.Category) error {
	res := db.connection.Save(&category)
	if res.Error != nil {
		res.Error = fmt.Errorf("gagal mengupdate kategori")
		return res.Error
	}
	return nil
}

func (db *categoryRepo) GetAllCategory(categoryGetAllRequest request.CategoryGetAllRequest) ([]database.Category, error) {
	var category []database.Category
	walletId, _ := strconv.ParseInt(categoryGetAllRequest.CategoryWalletId, 10, 64)
	err := db.connection.Where(database.Category{
		CategoryWalletID: walletId,
		CategoryType:     categoryGetAllRequest.CategoryType,
		CategoryIsActive: true,
	}).Offset((int(categoryGetAllRequest.CategoryPage) - 1) * 10).Limit(10).Find(&category)

	if err.Error != nil {
		err.Error = fmt.Errorf("gagal menemukan data")
		return nil, err.Error
	}
	return category, nil
}
