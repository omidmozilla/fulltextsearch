/*
	Component objects
	Product component
*/
package components

import (
	_ "database/sql"
	_ "encoding/json"
	_ "strconv"
	"log"
	_ "net/http"
	_ "context"
	_ "fmt"
	"github.com/gin-gonic/gin"

	//pb "rundoo.com/rpc"
)

type Product struct {
	Id        int    `json:"id"`
	Name  		string `json:"name"`
	Cateogry  string `json:"category"`
	SKU 			string `json:"sku"`
}

/*
	Creates a product
	@param product - Product object to add
*/
func CreateProduct(ctx *gin.Context)  {

	// product := &pb.Product{
	// 	Id:  "1",
	// 	Name: "First Product",
	// 	Category: "First Category",
	// 	Sku: "SKU",
	// }


}

/*
	Create user takes a user struct and inserts a user into the database
	@param searchText - String to search
*/
func SearchProducts(searchText string)  {
	log.Println("Searching products: ", searchText)
}

