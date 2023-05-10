export { AddProduct, SearchProducts }

async function AddProduct(product) {
  return fetch('http://localhost:8080/product', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(product)
  })
  .then(data => data.json())
  .catch( () => {
    throw Error("An error has occured during adding a product");
    }
  )
}

async function SearchProducts(searchTerm) {
  return fetch('http://localhost:8080/searchProducts', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(searchTerm)
  })
  .then(data => data.json())
  .catch( () => {
    throw Error("An error has occured during searching the products");
    }
  )
}