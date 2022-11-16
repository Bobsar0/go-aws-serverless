# Go-AWS Serverless Program

A simple serverless program in Go that implements CRUD operations. 

It features AWS serverless architecture stack: API Gateway, Lambda and DynamoDB database

## Getting Started

### Packages Version
Go v1.18.1

### Instructions

#### Local Build
This program is serverless meaning that it doesn't incorporate a server and thus cannot be run locally. It has to be built, zipped and uploaded to AWS lambda

- cd to `cmd` folder and build the main.go file using the command: `GOARCH=amd64 GOOS=linux go build main.go`
- zip up the main build
- Create an AWS go lambda instance and upload the zipped build with following parameters:
  - Runtime: `Go 1.x`
  - Hanlder name: `main`
  - Architecture: `x86_64` (at time of writing, only this architecture is supported for Go 1.x)
  - Minimum permissions role/policy template of `Simple microservice permissions (DynamoDB)`

- Create AWS API Gateway REST API instance, connect to the lambda and deploy

#### Local Test
Use curl, POSTMAN or equivalent API platform to test the following APIs:

- POST "{apiUrl}" (Create a new user)
body: {"email": "...", "firstName": "...", "lastName": "..."}
- GET "{apiUrl}" (Get all users)
- GET "{apiUrl}?email=..." (Get a user by email)
- PUT "{apiUrl}" (Update a user)
body: {"email": "...", "firstName": "...", "lastName": "..."}
- DELETE "{apiUrl}?email=..." (Delete a user by email)