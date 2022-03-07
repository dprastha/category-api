package test

import (
	"belajar-golang-rest-api/app"
	"belajar-golang-rest-api/model/domain"
	"belajar-golang-rest-api/repository"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func truncateProduct(db *sql.DB) {
	db.Exec("TRUNCATE product")
}

func creteDumpDataProduct(dataLength int) (*sql.DB, []domain.Product) {
	db := app.SetupTestDB()
	truncateProduct(db)

	var products []domain.Product

	i := 1
	for i <= dataLength {
		tx, _ := db.Begin()
		productRepository := repository.NewProductRepository()
		product := productRepository.Save(context.Background(), tx, domain.Product{
			Name: "Product Test" + strconv.Itoa(i),
		})
		tx.Commit()

		products = append(products, product)
		i++
	}

	return db, products

}

func TestCreateProductSuccess(t *testing.T) {
	router := SetupTest()

	requestBody := strings.NewReader(`{"name": "Macbook Pro"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/products", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	reponse := recorder.Result()
	assert.Equal(t, 200, reponse.StatusCode)

	body, _ := io.ReadAll(reponse.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Macbook Pro", responseBody["data"].(map[string]interface{})["name"])
}
func TestCreateProductFailed(t *testing.T) {
	router := SetupTest()

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/products", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateProductSuccess(t *testing.T) {
	db, product := creteDumpDataProduct(1)

	router := app.SetupRouter(db)

	requestBody := strings.NewReader(`{"name": "Product Test1"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/products/"+strconv.Itoa(product[0].Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, product[0].Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, product[0].Name, responseBody["data"].(map[string]interface{})["name"])
}
func TestUpdateProductFailed(t *testing.T) {
	db, product := creteDumpDataProduct(1)

	router := app.SetupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/products/"+strconv.Itoa(product[0].Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}
func TestDeleteProductSuccess(t *testing.T) {
	db, product := creteDumpDataProduct(1)

	router := app.SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/products/"+strconv.Itoa(product[0].Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var resposneBody map[string]interface{}
	json.Unmarshal(body, &resposneBody)

	assert.Equal(t, http.StatusOK, int(resposneBody["code"].(float64)))
	assert.Equal(t, "OK", resposneBody["status"])
}
func TestDeleteProductFailed(t *testing.T) {
	db, _ := creteDumpDataProduct(1)

	router := app.SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/products/999", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}
func TestGetProductSuccess(t *testing.T) {
	db, product := creteDumpDataProduct(1)

	router := app.SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/products/"+strconv.Itoa(product[0].Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, product[0].Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, product[0].Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetProductFailed(t *testing.T) {
	db, _ := creteDumpDataProduct(1)

	router := app.SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/products/999", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListProductSuccess(t *testing.T) {
	db, product := creteDumpDataProduct(2)

	router := app.SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/products", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var products = responseBody["data"].([]interface{})

	productResponse1 := products[0].(map[string]interface{})
	productResponse2 := products[1].(map[string]interface{})

	assert.Equal(t, product[0].Id, int(productResponse1["id"].(float64)))
	assert.Equal(t, product[0].Name, productResponse1["name"])

	assert.Equal(t, product[1].Id, int(productResponse2["id"].(float64)))
	assert.Equal(t, product[1].Name, productResponse2["name"])
}
func TestListProductFailed(t *testing.T) {
	db, _ := creteDumpDataProduct(2)

	router := app.SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/products", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
