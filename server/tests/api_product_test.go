package main

/*

	Testing File for the api/product api

	Runs create and valdiation tests

*/

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"rundoo.com/server/apis"
	"testing"
	_ "net/url"
	"os"
	"log"
	"bytes"
)

func createProduct(t *testing.T) {
	mockResponse := `{"success":"New product has been created"}`
	body := []byte(`{
		"name": "test name",
		"category": "test category",
		"sku": "test sku"
	}`)

	r := setUpRouter()
	r.POST("/product", apis.CreateProduct)
	log.Println("Body is ", bytes.NewBuffer(body))
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func createProductValidation(t *testing.T) {
	bodies := map[string]string {
		"name" : string([]byte(`{
			"name": "",
			"category": "",
			"sku": ""
		}`)),
		"category" : string([]byte(`{
			"name": "test name",
			"category": "",
			"sku": ""
		}`)),
		"sku" : string([]byte(`{
			"name": "test name",
			"category": "test category",
			"sku": ""
		}`)),
	}

	responses := map[string]string {
		"name" : "Product name must not be empty",
		"category" : "Product category must not be empty",
		"sku" : "Product sku must not be empty",
	}

	r := setUpRouter()
	r.POST("/product", apis.CreateProduct)
	for index, value := range(bodies) {			
		req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer([]byte(value)))
		req.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		assert.Equal(t, `{"failure":"` + responses[index] + `"}`, string(responseData))
		assert.Equal(t, http.StatusOK, w.Code)
	}
}

func TestProductsAPI(t *testing.T) {
	file, err := os.OpenFile("../logs/app_test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	defer file.Close()

	//setup()
	createProduct(t)
	createProductValidation(t)
}
