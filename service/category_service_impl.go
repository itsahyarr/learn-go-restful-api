package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/itsahyarr/learn-go-restful-api/exception"
	"github.com/itsahyarr/learn-go-restful-api/helper"
	"github.com/itsahyarr/learn-go-restful-api/model/domain"
	"github.com/itsahyarr/learn-go-restful-api/model/web"
	"github.com/itsahyarr/learn-go-restful-api/repository"
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

// ======================================================================================================================
// CREATE / SAVE DATA
// ======================================================================================================================
func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

// ======================================================================================================================
// UPDATE DATA
// ======================================================================================================================
func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	// helper.PanicIfError(err)
	// HANDLE WITH EXCEPTION HANDLER
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Set Name
	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

// ======================================================================================================================
// DELETE DATA
// ======================================================================================================================
func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	// helper.PanicIfError(err)
	// HANDLE WITH EXCEPTION HANDLER
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

// ======================================================================================================================
// FIND SINGLE DATA BY ID
// ======================================================================================================================
func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	// helper.PanicIfError(err)
	// HANDLE WITH EXCEPTION HANDLER
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

// ======================================================================================================================
// FIND ALL DATA
// ======================================================================================================================
func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
