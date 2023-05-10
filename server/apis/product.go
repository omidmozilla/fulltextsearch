package apis

/*
	Product API
	creates and searches products 
*/

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"context"
	"log"
	"fmt"
	"rundoo.com/rpc"
	"rundoo.com/config"
)


type APIError struct {
	Message string
	Err     error
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

type MyProduct struct {
	apiProduct rpc.Product
}


func (product *MyProduct) validate() error {

	if len(product.apiProduct.Name) == 0 {
		return &APIError{"Product name must not be empty", nil}
	}
	if len(product.apiProduct.Category) == 0 {
		return &APIError{"Product category must not be empty", nil}
	}
	if len(product.apiProduct.Sku) == 0 {
		return &APIError{"Product sku must not be empty", nil}
	}
	
	return nil
}

func getRPCClientURL() string {
	rpcConfig := config.GetRPCConfig()
	return rpcConfig.Host + ":" + rpcConfig.Port
}

/*
	CreateProduct - calls the add product rpc and sets the success or failure HTTP response of the context 
	@param ctx - gin.Context pointer contining the request context
*/
func CreateProduct(ctx *gin.Context) {

	rpcProduct := rpc.Product{}
	if err := ctx.BindJSON(&rpcProduct); err != nil {
		log.Fatal("An unrecoverable error has occured on binding the body for create product %s: ", err)
		return
	}

	product := MyProduct{apiProduct: rpcProduct}

	valid := product.validate()
	if valid != nil {
		log.Println("Error has occured validating create product ", valid)
		ctx.JSON(http.StatusOK, gin.H{"failure": valid.Error()})
		return
	}

	
	log.Println("Product to add: ", product)
	client := rpc.NewProductServiceProtobufClient(config.GetRPCURL(), &http.Client{})
	resp, err := client.AddProduct(context.Background(), &rpc.AddProductReq{Product: &product.apiProduct})
	if err == nil {
		log.Println("Product created success: ", resp)
		ctx.JSON(http.StatusOK, gin.H{"success": "New product has been created"})
	} else {
		log.Println("Product created failure: ", err)
		ctx.JSON(http.StatusOK, gin.H{"failure": "An error has occured"})
	}
}

/*
	SearchProducts - calls the search product rpc and sets the success or failure HTTP response of the context 
	@param ctx - gin.Context pointer contining the request context
*/
func SearchProducts(ctx *gin.Context) {

	searchTerm := &rpc.Search{}
	if err := ctx.BindJSON(&searchTerm); err != nil {
		log.Fatal("An unrecoverable error has occured on binding the body for search term %s: ", err)
		return
	}

	log.Println("Product to search: ", searchTerm)
	client := rpc.NewProductServiceProtobufClient(config.GetRPCURL(), &http.Client{})
	resp, err := client.SearchProducts(context.Background(), &rpc.SearchProductReq{SearchTerm: searchTerm})
	if err == nil {
		log.Println("Search products success: ", resp)
		ctx.JSON(http.StatusOK, gin.H{"success": resp})
	} else {
		log.Println("Product created failure: ", err)
		ctx.JSON(http.StatusOK, gin.H{"failure": "An error has occured"})
	}
}