# Products Service 
This is a simple service that provisions an endpoint that retrieves products.  

## Overview 
`/products` endpoint retrieves products from the database (with filtering applied), discounts applied (if necessary) and the data returned as JSON object. 

### Important notes
* The `db` package declares a type that any DBMS, im-memory storage, etc must implement in order to retrieve products. For this project, I have `mongo.go` that shows how products would be retrieved from a mongo database. It's implementation is hidden for this project.

* The `workflow` package holds the business logic for this app. It also declares a type that must be implemented. 

## Usage
Clone project and `cd` into project foler

### Starting server
``` bash
$ make run
```  

### Running Tests
``` bash
$ make test
```




