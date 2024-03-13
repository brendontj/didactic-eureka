# Didactic Eureka Project

This project is a Golang API Rest to handle CRUD operations related to customers using RabbitMQ and PostgresQL.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Features](#features)
- [Operations](#operations)
- [Implementations and further improvements](#implementations-and-further-improvements)

## Requirements

- Go 1.21 or later
- Docker
- Docker Compose
- Additional dependencies as specified in `go.mod`

## Installation

1. Clone this repository:

   ```sh
   git clone https://github.com/brendontj/didactic-eureka.git 
   cd didactic-eureka
   ```

2. Install dependencies:

   ```sh
   go mod download
   ```

3. Copy the .env.example file to .env and fill in the environment variables:

   ```sh
   cp .env.example .env
   ```

4. Start docker container with rabbitmq and postgresql dependencies:
    ```sh
    docker-compose up -d 
    ```

5. Run all migrations:

   ```sh
   make migrate 
   ```

6. Run the core application:

   ```sh
   go run cmd/api/main.go
   ```

7. Run the worker:

   ```sh
   go run worker/main.go
   ```

8. Send requests to the API at the following address:
    - localhost:8080

## Operations

### Create a new customer

```sh
curl --request POST \
  --url http://localhost:8080/api/v1/customers \
  --header 'Content-Type: application/json' \ 
  --data '{
	"name": "testName",
	"email": "test@gmail.com",
	"phone": "123456789",
	"birth_date": "1995-06-08",
	"document": "123456789",
	"address": {
		"street": "aStreet",
		"number": "aNumber",
		"zip_code": "123456789",
		"neighborhood": "aNeighborhood",
		"city": "aCity",
		"state": "aState",
		"country": "aCountry",
		"complement": "aComplement"
	}
}'
```

### Find all customers

```sh
curl --request GET \
  --url http://localhost:8080/api/v1/customers \
  --header 'Content-Type: application/json' 
```

### Find customer by ID

```sh
curl --request GET \
  --url http://localhost:8080/api/v1/customers/{customerId} \
  --header 'Content-Type: application/json' 
 ```

### Update customer

```sh
curl --request PUT \
  --url http://localhost:8080/api/v1/customers/{customerId}/version/{customerVersion} \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "newName",
	"email": "newEmail@gmail.com",
	"phone": "33333333",
	"birth_date": "1995-06-20",
	"document": "1345679843",
	"address": {
		"street": "aStreet",
		"number": "123",
		"zip_code": "123325432",
		"neighborhood": "aNeighborhood",
		"city": "aCity",
		"state": "aState",
		"country": "aCountry",
		"complement": "aComplement"
	}
}'
```

### Delete customer by ID

```sh
curl --request DELETE \
  --url http://localhost:8080/api/v1/customers/{customerId} \
  --header 'Content-Type: application/json' 
```

## Implementations and further improvements

- Create a new customer 
- Find all customers
- Find customer by ID
- Update customer by ID
- Delete customer by ID
- (not implemented) Add logging to the application
- (not implemented) Pagination in the list of customers
- (not implemented) Validation of the domain fields
- (not implemented) Unit and integrations tests 
- (not implemented) Api swagger documentation


