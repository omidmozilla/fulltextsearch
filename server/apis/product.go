package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"context"
	"log"
	"rundoo.com/rpc"
)

func CreateProduct(ctx *gin.Context) {

	product := &rpc.Product{}
	if err := ctx.BindJSON(&product); err != nil {
		log.Fatal("An unrecoverable error has occured on binding the body for create product %s: ", err)
		return
	}

	log.Println("Product to add: ", product)
	client := rpc.NewProductServiceProtobufClient("http://localhost:8088", &http.Client{})
	resp, err := client.AddProduct(context.Background(), &rpc.AddProductReq{Product: product})
	if err == nil {
		log.Println("Product created success: ", resp)
		ctx.JSON(http.StatusOK, gin.H{"success": "New product has been created"})
	} else {
		log.Println("Product created failure: ", err)
		ctx.JSON(http.StatusOK, gin.H{"failure": "An error has occured"})
	}
}

func SearchProducts(ctx *gin.Context) {

	searchTerm := &rpc.Search{}
	if err := ctx.BindJSON(&searchTerm); err != nil {
		log.Fatal("An unrecoverable error has occured on binding the body for search term %s: ", err)
		return
	}

	log.Println("Product to search: ", searchTerm)
	client := rpc.NewProductServiceProtobufClient("http://localhost:8088", &http.Client{})
	resp, err := client.SearchProducts(context.Background(), &rpc.SearchProductReq{SearchTerm: searchTerm})
	if err == nil {
		log.Println("Search products success: ", resp)
		ctx.JSON(http.StatusOK, gin.H{"success": resp})
	} else {
		log.Println("Product created failure: ", err)
		ctx.JSON(http.StatusOK, gin.H{"failure": "An error has occured"})
	}
}