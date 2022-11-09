package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/bobsar0/go-aws-serverless/pkg/user"
	"github.com/bobsar0/go-aws-serverless/user"

	"github.com/aws/aws-sdk-go/service/lambda"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI)  (*events.APIGatewayProxyResponse, error){
	email := req.QueryStringParameters["email"]

	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynamoClient)

		if err!=nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error)})
		}
		return apiResponse(http.StatusOK, result)
	}

	result, err := user.FetchUsers(tableName, dynamoClient)

		if err!=nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error)})
		}
		return apiResponse(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI)  {
	
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI)  {
	
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI)  {
  
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error)  {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}