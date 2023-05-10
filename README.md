# Full Text Search

Full Text Search is an example program that allows you to add products and performs a full text search

## Technologies

Project is created with:

- client - React Client
- rpc - Twirp on protobuf rpc library
- server - Go API server
- storage - Go MongoDB server
- config - Configuration files

## Installation

You will need mongodb installed and run the following commands:\
use fulltextsearch\
db.products.createIndex({ name: "text", category: "text", sku: "text" })

Change config/config.go for any custom ports

## Usage

Start storage server\
cd storage\
go run storage.go

Start backend server\
cd server\
go run server.go

Start client\
cd client
npm start

Browse to http://localhost:3000/

## Tests

To run the test start the storage server

cd server/tests\
go test

## Contributing

By Omid Mozilla

## License
