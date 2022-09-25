package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/ramadoiranedar/go_restapi/exception"
	"github.com/ramadoiranedar/go_restapi/helper"
	"github.com/ramadoiranedar/go_restapi/model/domain"
	"github.com/ramadoiranedar/go_restapi/model/web"
	"github.com/ramadoiranedar/go_restapi/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	category := domain.Category{
		Name: request.Name,
	}
	category = service.CategoryRepository.Create(ctx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	category, err := service.CategoryRepository.FindById(ctx, request.Id)
	if err != nil {
		helper.PanicIfError(err)
	}
	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.CategoryRepository.Delete(ctx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	categories := service.CategoryRepository.FindAll(ctx)
	return helper.ToCategoryResponses(categories)
}
