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
		"rundoo.com/config"
)

func connectMongoClient(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
			log.Fatal("Unable to establish a connect to mongo :", err)			
	}

	return client, err
}

type ProductServiceServer struct{}
const PRODUCT_COLLECTION = "products"

func (s *ProductServiceServer) AddProduct(ctx context.Context, req *rpc.AddProductReq) (*rpc.AddProductResp, error) {

	log.Println("Product to add: ", req)
	client, err := connectMongoClient(ctx)

	dbConfig := config.GetDBConfig()
	collection := client.Database(dbConfig.Name).Collection("products")
	doc := bson.D{
			{"name", req.Product.Name},
			{"category", req.Product.Category},
			{"sku", req.Product.Sku},
	}
	result, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		log.Fatal("Unable to insert a collection :", err)	
	}

	log.Printf("Inserted document with ID %v\n", result.InsertedID)
	return &rpc.AddProductResp{Success: true}, nil
}

func (s *ProductServiceServer) SearchProducts(ctx context.Context, req *rpc.SearchProductReq) (*rpc.SearchProductResp, error) {

	log.Println("Product to search: ", req)
	client, err := connectMongoClient(ctx)

	dbConfig := config.GetDBConfig()
	log.Println("CONFIG: ", dbConfig.Name)
	collection := client.Database(dbConfig.Name).Collection("products")
	search := bson.M{
    "$text": bson.M{
    "$search": req.SearchTerm.SearchTerm,
   },
	}

	cursor, err := collection.Find(context.Background(), search)
	products := []*rpc.Product{}
	if err != nil { 
		log.Fatal("Unable to complete the find on a collection :", err)	
		defer cursor.Close(ctx)
	}

	for cursor.Next(ctx) {	
			var result bson.M			
			product := rpc.Product{}
			err := cursor.Decode(&result)
			if err != nil {
				log.Fatal("Unable to decode on a collection :", err)
			} else {
				result["id"] = result["_id"].(primitive.ObjectID).Hex()
				bsonBytes, _ := bson.Marshal(result)
				bson.Unmarshal(bsonBytes, &product)
				products = append(products, &product)
			}
	}
	
	log.Printf("Search completed: %v\n", products)
	return &rpc.SearchProductResp{Products: products}, nil
}

// Run the implementation in a local server
func main() {

		file, err := os.OpenFile("logs/storage.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Unable to open the log file for storage: ", err)
		}
		defer file.Close()
		log.SetOutput(file)

		log.Println("Storage server running")
    productHandler := rpc.NewProductServiceServer(&ProductServiceServer{})
    mux := http.NewServeMux()
    mux.Handle(productHandler.PathPrefix(), productHandler)
    http.ListenAndServe(":8088", mux)
}
