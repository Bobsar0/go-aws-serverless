package user

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/bobsar0/go-aws-serverless/pkg/validators"
)

var (
	ErrorFailedToFetchRecord     = "Failed to fetch record"
	ErrorFailedToUnmarshalRecord = "Failed to unmarshal record"
	ErrorInvalidUserData = "Invalid user data"
	ErrorInvalidEmail = "Invalid email"
	ErrorCouldNotMarshalItem = "Could not marshal item"
	ErrorCouldNotDeleteItem = "Could not delete item"
	ErrorDynamoCouldNotPutItem = "Dynamo could not put items into"
	ErrorUserAlreadyExists = "User already exists"
	ErrorUserDoesNotExist = "User does not exist"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func FetchUser(email string, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	result, err := dynamoClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	user := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, user)

	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	return user, nil
}

func FetchUsers(tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*[]User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dynamoClient.Scan(input)

	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	users := new([]User)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, users)

	return users, nil

}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (
	*User, error,
	) {
		var user User

		if err := json.Unmarshal([]byte(req.Body), &user); err != nil {
			return nil, errors.New(ErrorInvalidUserData)
		}

		if !validators.IsEmailValid(user.Email) {
			return nil, errors.New(ErrorInvalidEmail)
		}

		currUser, _ := FetchUser(user.Email, tableName, dynamoClient)

		if currUser != nil && len(currUser.Email) != 0 {
			return nil, errors.New(ErrorUserAlreadyExists)
		}

		av, err := dynamodbattribute.MarshalMap(user)
		if err != nil {
			return nil, errors.New(ErrorCouldNotMarshalItem)
		}

		input:= &dynamodb.PutItemInput{
			Item: av,
			TableName: aws.String(tableName),
		}

		_, err = dynamoClient.PutItem(input)
		if err != nil {
			return nil, errors.New(ErrorDynamoCouldNotPutItem)
		}

		return &user, nil
}

func UpdateUser(email string, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) {

}

func DeleteUser() error {

}

func UnhandledMethod() {

}
