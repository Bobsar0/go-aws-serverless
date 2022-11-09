package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)


func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}

	res.StatusCode = status

	stringBody, _ := json.Marshal(body)
	res.Body = string(stringBody)

	return &res, nil
}