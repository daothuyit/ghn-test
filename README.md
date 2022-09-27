
# ghn-test
A REST API with Golang, MongoDB and Echo Framework.



## Features

Product Management with CRUD APIs.



## Installation

Run file build.sh

```bash
  ./build.sh
```
    
## Usage

Get JWT to authenticate
```javascript
curl --location --request GET 'http://localhost:17001/jwt'
```
```javascript
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQyNDkzNTF9.B-QMXFtNDckC3j6sMwbhUEnt23ER5pV51tCBd5nmAHk"
}
```

After that, using this "access_token" to call CRUD APIs:

API Create a new product
```javascript
curl --location --request POST 'http://localhost:17001/products' \
--header 'access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQyNDkzNTF9.B-QMXFtNDckC3j6sMwbhUEnt23ER5pV51tCBd5nmAHk' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "product 1",
    "description": "description of product 1",
    "price": 45000
}'
```

API Update a new product
```javascript
curl --location --request PUT 'http://localhost:17001/products/63326166620e64247d45d601' \
--header 'access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQyNDkzNTF9.B-QMXFtNDckC3j6sMwbhUEnt23ER5pV51tCBd5nmAHk' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "product 1",
    "description": "description of product 1",
    "price": 55000
}'
```

API Get a product by Id
```javascript
curl --location --request GET 'http://localhost:17001/products/63326166620e64247d45d601' \
--header 'access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQyNDkzNTF9.B-QMXFtNDckC3j6sMwbhUEnt23ER5pV51tCBd5nmAHk'
```

API Get all products
```javascript
curl --location --request GET 'http://localhost:17001/products?max=10&offset=1' \
--header 'access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQyNDkzNTF9.B-QMXFtNDckC3j6sMwbhUEnt23ER5pV51tCBd5nmAHk'
```

API delete a product by Id
```javascript
curl --location --request DELETE 'http://localhost:17001/products/63326166620e64247d45d601' \
--header 'access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQyNDkzNTF9.B-QMXFtNDckC3j6sMwbhUEnt23ER5pV51tCBd5nmAHk'
```
## Related

Postman Document: https://documenter.getpostman.com/view/1876235/2s83YeAMB5

