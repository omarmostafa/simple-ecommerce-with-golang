#Simple E-Commerce Rest APIs

This is an example of simple e-commerce rest APIS with clean architecture, using dependency injection , design patterns and following Solid principle


## install 
Clone project 

```bash
git clone https://github.com/omarmostafa/simple-ecommerce-with-golang.git
``` 

Setup Dependency

```bash
        go get github.com/Jeffail/gabs/v2 v2.1.0
	go get github.com/dgrijalva/jwt-go v3.2.0+incompatible
	go get github.com/go-chi/chi v4.0.2+incompatible
	go get github.com/go-playground/locales v0.12.1 // indirect
	go get github.com/go-playground/universal-translator v0.16.0 // indirect
	go get github.com/go-sql-driver/mysql v1.4.1
	go get github.com/gorilla/mux v1.7.3
	go get github.com/jinzhu/gorm v1.9.10
	go get github.com/joho/godotenv v1.3.0
	go get github.com/leodido/go-urn v1.1.0 // indirect
	go get github.com/lib/pq v1.1.1
	go get github.com/stretchr/testify v1.4.0 // indirect
	go get golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c
	go get gopkg.in/go-playground/validator.v9 v9.29.1
	go get gopkg.in/matryer/respond.v1 v1.0.1
	go get gopkg.in/validator.v2 v2.0.0-20180514200540-135c24b11c19 // indirect``` 
```
 Run APP 
 ```bash
go build && ./simple-ecommerce
 ``` 
 
 ##Intro
 
 This is an example for writing clean code in go to follow SOLID principle 
 
 The aim of this pattern is to decoupled systems that the implementation of lower level domain is not a concern of the implementor and all frameworks are independent 
 
 - Independent of database , business logic should be isolated to the database , the system should be swapped between database management systems without affect in business logic 
 - Independent of 3rd party and external library , business logic should be isolated to any 3rd party because if there any change of this it shouldn't affect business logic 
 
 ## Controllers
 
 controllers is the handler of request and responses , no business logic or data access layer should be done in controller it should be separated 
 
 ## Models 
 
 Models is the data object from/to database 
 
 ## Repositories 
 
repositories folder hosts all structs under repositories namespace, repositories is where the implementation of data access layer. All queries and data operation from / to database should happen here, and the implementor should be agnostic of what is the database engine is used, how the queries is done, all they care is they can pull the data according to the interface they are implementing.
 
 ## Services 
 
 services folder hosts all structs under services namespace, services is where the business logic lies on, it handles controller request and fetch data from data layer it needs and run their logic to satisfy what controller expect the service to return.
 
 controller might implement many services interface to satisfy the request needs, and controller should be agnostic of how services implements their logic, all they care is that they should be able to pull the result they need according to the interface they implements.
 
 ## Transformers
 
 Transformers is used to transform data from model and convert it to right format 
 
 ## Middleware
 Middleware provide a convenient mechanism for filtering HTTP requests entering your application
, such as authentication or authorization 

## Authentication using JWT

we use JWT to generate tokens for each authenticated tokens
,this token would be included in the header of the subsequent request made to the API server, this method allow us to identify every users that make calls to our API
 