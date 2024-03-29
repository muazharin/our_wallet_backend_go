package services

import (
	"strconv"
	"time"

	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"github.com/muazharin/our_wallet_backend_go/src/entities/request"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
)

type CategoryService interface {
	AddCategory(categoryAddRequest request.CategoryAddRequest) error
	UpdateCategory(categoryUpdateRequest request.CategoryUpdateRequest) error
	DeleteCategory(categoryDeleteRequest request.CategoryDeleteRequest) error
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

func (s *categoryService) AddCategory(categoryAddRequest request.CategoryAddRequest) error {
	var category database.Category
	category.CategoryID = time.Now().Unix()
	category.CategoryWalletID, _ = strconv.ParseInt(categoryAddRequest.CategoryWalletID, 10, 64)
	category.CategoryTitle = categoryAddRequest.CategoryTitle
	category.CategoryType = categoryAddRequest.CategoryType
	category.CategoryIsActive = true
	err := s.categoryRepo.AddCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *categoryService) UpdateCategory(categoryUpdateRequest request.CategoryUpdateRequest) error {

	res, err := s.categoryRepo.GetCategoryById(categoryUpdateRequest.CategoryID)
	if err != nil {
		return err
	}
	res.CategoryTitle = categoryUpdateRequest.CategoryTitle
	err = s.categoryRepo.UpdateCategory(res)
	if err != nil {
		return err
	}
	return nil
}

func (s *categoryService) DeleteCategory(categoryDeleteRequest request.CategoryDeleteRequest) error {
	res, err := s.categoryRepo.GetCategoryById(categoryDeleteRequest.CategoryID)
	if err != nil {
		return err
	}
	res.CategoryIsActive = false
	err = s.categoryRepo.DeleteCategory(res)
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
