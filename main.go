package main

import (
	"os"
	"github.com/bobsar0/go-aws-serverless/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
		"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/aws/aws-sdk-go/service/lambda"
)

var dynamoClient dynamodbiface.DynamoDBAPI

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return
	}

	dynamoClient := dynamodb.New(awsSession)
	lambda.Start(handler)
}

const tableName = "LambdaInGoUser"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
		case "GET":
    return handlers.GetUser(req, tableName, dynamoClient)
		case "POST":
			return handlers.CreateUser(req, tableName, dynamoClient)
		case "PUT":
			return handlers.UpdateUser(req, tableName, dynamoClient)
		case "DELETE":
			return handlers.DeleteUser(req, tableName, dynamoClient)
			default:
      return handlers.UnhandledMethod()
	}
}