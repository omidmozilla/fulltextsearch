package main

import (
    "context"
    "net/http"
		"log"
		"os"
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/mongo/options"
		"go.mongodb.org/mongo-driver/bson/primitive"
    "rundoo.com/rpc"
)

type ProductServiceServer struct{}

func (s *ProductServiceServer) AddProduct(ctx context.Context, req *rpc.AddProductReq) (*rpc.AddProductResp, error) {

	log.Println("Product to add: ", req)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
			log.Println(err)
	}

	collection := client.Database("rundoo").Collection("products")
	doc := bson.D{
			{"name", req.Product.Name},
			{"category", req.Product.Category},
			{"sku", req.Product.Sku},
	}
	result, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
			// handle error
	}

	log.Printf("Inserted document with ID %v\n", result.InsertedID)
	return &rpc.AddProductResp{Success: true}, nil
}

func (s *ProductServiceServer) SearchProducts(ctx context.Context, req *rpc.SearchProductReq) (*rpc.SearchProductResp, error) {

	log.Println("Product to search: ", req)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
			log.Println(err)
	}

	collection := client.Database("rundoo").Collection("products")

	search := bson.M{
    "$text": bson.M{
    "$search": req.SearchTerm.SearchTerm,
   },
	}

	cursor, err := collection.Find(context.Background(), search)
	// if err != nil { // Handle the error
	// 		// handle error
	// }

	products := []*rpc.Product{}
	 // Find() method raised an error
	if err != nil {
		log.Println("Finding all products ERROR:", err)
		defer cursor.Close(ctx)
	} else {
		// iterate over docs using Next()
		for cursor.Next(ctx) {
				// declare a result BSON object	
				var result bson.M			
				product := rpc.Product{}
				err := cursor.Decode(&result)
				// If there is a cursor.Decode error
				if err != nil {
					// Handle the error log.Println("cursor.Next() error:", err)
				} else {
					result["id"] = result["_id"].(primitive.ObjectID).Hex()
					bsonBytes, _ := bson.Marshal(result)
					bson.Unmarshal(bsonBytes, &product)
					products = append(products, &product)
				}
		}
	}	
	
	log.Printf("Search completed: %v\n", products)
	return &rpc.SearchProductResp{Products: products}, nil
}

// Run the implementation in a local server
func main() {

		file, err := os.OpenFile("logs/storage.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		log.SetOutput(file)

		log.Println("Storage server running")

    productHandler := rpc.NewProductServiceServer(&ProductServiceServer{})
    mux := http.NewServeMux()
    mux.Handle(productHandler.PathPrefix(), productHandler)
    http.ListenAndServe(":8088", mux)
}
