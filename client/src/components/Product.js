import React, { useState } from 'react';
import Grid from '@mui/material/Grid'
import { Typography, TextField, Button } from '@mui/material';
import { IsEmptyString } from '../utils/data.utils.js'
import { AddProduct, SearchProducts } from '../data/Product.js'

export { Product, ProductSearch }

function Product() {
  const [productName,       setProductName]     = useState('');
  const [productCategory,   setProductCategory] = useState('');
  const [productSku,        setProductSku]      = useState('');
  const [message,           setMessage]         = useState('');
  const [productNameMsg,    setProductNameMsg]  = useState('');

  const isValid = (productName) => {

    var returnValue = null
    if( IsEmptyString(productName) ) {
      setProductNameMsg("Product Name must not be empty")
      returnValue = false
    }

    return returnValue
  }

  const handleSubmit = async e => {
    e.preventDefault();

    setMessage('')
    setProductNameMsg('')
    var isValidMessage = isValid(productName);
    if( isValidMessage != null ) {
      setMessage(isValidMessage)
    } else {
      try {
        const _ = await AddProduct({name: productName, category: productCategory, sku: productSku});
        // if (response.hasOwnProperty('success')) {   
        // } else if (response.hasOwnProperty('failure')) {
        // } else {
        // }
      }
      catch (error) {
        //setMessage(error.message)
      }
    }
  }

  return (
    <Grid container rowSpacing={0} columnSpacing={4} sx={{padding: 10}}>        
        <Typography component="h1" variant="h5">
          {productNameMsg}
        </Typography>
        <Typography component="h1" variant="h5">
          {message}
        </Typography>            
        <form noValidate onSubmit={handleSubmit}>
          <Grid item xs={12}>              
              <TextField
                variant="outlined"
                margin="normal"
                required                  
                id="productName"
                name="productName"
                // error={productName == '' ? false : true}
                label="Product Name"
                onChange={e => setProductName(e.target.value)}
              />
              <TextField
                variant="outlined"
                margin="normal"
                required                  
                id="productCategory"
                name="productCategory"
                // error={productName == '' ? false : true}
                label="Product Category"
                onChange={e => setProductCategory(e.target.value)}
              />      
              <TextField
                variant="outlined"
                margin="normal"
                required                  
                id="productSku"
                name="productSku"
                // error={productName == '' ? false : true}
                label="Product Sku"
                onChange={e => setProductSku(e.target.value)}
              />                        
          </Grid>
          <Grid item xs={4} md={10}>     
            <Button
              type="submit"                  
              variant="contained"
              color="primary"
            >
              Add Product
            </Button>              
        </Grid>
      </form>
    </Grid>
  );
}

class ProductRow extends React.Component {
  render() {
    const product = this.props.product;
    const name = product.name
    const category = product.category
    const sku = product.sku

    return (
      <tr>
        <td>{name}</td>
        <td>{category}</td>
        <td>{sku}</td>        
      </tr>
    );
  }
}

class ProductTable extends React.Component {
  render() {
    const rows = [];
    this.props.products.forEach((product) => {     
      rows.push(
        <ProductRow
          product={product}
          key={product.id} 
        />
      );
    });

    return (
      <div>
      <Typography>Search Results</Typography>      
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th>Category</th>
              <th>SKU</th>
            </tr>
          </thead>
          <tbody>{rows}</tbody>
        </table>
      </div>
    );
  }
}

function ProductSearch() {
  const [search,    setSearch]     = useState('');
  const [products,  setProducts]   = useState([]);
  const [message,   setMessage]   = useState([]);

  const handleSubmit = async e => {
    e.preventDefault();

      setMessage('')
      try {
        const response = await SearchProducts({searchTerm: search});
        if (response.hasOwnProperty('success')) {
          if (response.success.hasOwnProperty('products')) {
            setProducts(response.success.products)
          }
          else {
            setMessage('No results found')
            setProducts([])
          }
        }
      }
      catch (error) {
        setMessage(error.message)
      }
  }
  
  return (
    <div>
    <form noValidate onSubmit={handleSubmit}>      
      <Grid container rowSpacing={2} columnSpacing={2} sx={{padding: 10}}>
        <Grid item xs={2}> 
          <Typography>{message}</Typography>
          <TextField
                variant="outlined"
                margin="normal"
                required                  
                id="search"
                name="search"
                // error={productName == '' ? false : true}
                label="Search"
                onChange={e => setSearch(e.target.value)}
          />   
        </Grid>   
        <Grid item xs={12}> 
          <Button
              type="submit"                  
              variant="contained"
              color="primary"
            >
              Search
          </Button> 
        </Grid>              
        <Grid item xs={2}> 
          <ProductTable 
            products={products} 
            filterText={search} 
        />
        </Grid>
      </Grid>
    </form>
    </div>
  );
}