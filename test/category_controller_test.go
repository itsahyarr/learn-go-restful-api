package test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/itsahyarr/learn-go-restful-api/app"
	"github.com/itsahyarr/learn-go-restful-api/controller"
	"github.com/itsahyarr/learn-go-restful-api/helper"
	"github.com/itsahyarr/learn-go-restful-api/middleware"
	"github.com/itsahyarr/learn-go-restful-api/repository"
	"github.com/itsahyarr/learn-go-restful-api/service"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:thinkpad@tcp(172.20.0.3)/belajar_golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	// db := setupTestDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

// Menghapus data terlebih dahulu
func truncateDB(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

// ======================================================================================
// TESTING SCENARIO
// ======================================================================================
func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}
func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
	// assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}
func TestUpdateCategorySuccess(t *testing.T) {}
func TestUpdateCategoryFailed(t *testing.T)  {}
func TestGetCategorySuccess(t *testing.T)    {}
func TestGetCategoryFailed(t *testing.T)     {}
func TestDeleteCategorySuccess(t *testing.T) {}
func TestDeleteCategoryFailed(t *testing.T)  {}
func TestListCategoriesSuccess(t *testing.T) {}
func TestUnauthorized(t *testing.T)          {}
