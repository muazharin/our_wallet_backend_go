package services

import (
	"strconv"
	"time"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type CategoryService interface {
	AddCategory(categoryAddRequest request.CategoryAddRequest, userId int64) error
	GetAllCategory(categoryGetAllRequest request.CategoryGetAllRequest) ([]database.Category, error)
}

type categoryService struct {
	categoryRepo repositories.CategoryRepo
}

func NewCategoryService(categoryRepo repositories.CategoryRepo) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) AddCategory(categoryAddRequest request.CategoryAddRequest, userId int64) error {
	var category database.Category
	category.CategoryID = time.Now().Unix()
	category.CategoryUserID = userId
	category.CategoryWalletID, _ = strconv.ParseInt(categoryAddRequest.CategoryWalletID, 10, 64)
	category.CategoryTitle = categoryAddRequest.CategoryTitle
	category.CategoryType = categoryAddRequest.CategoryType
	err := s.categoryRepo.AddCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *categoryService) GetAllCategory(categoryGetAllRequest request.CategoryGetAllRequest) ([]database.Category, error) {
	res, err := s.categoryRepo.GetAllCategory(categoryGetAllRequest)
	if err != nil {
		return []database.Category{}, nil
	}
	return res, nil
}
