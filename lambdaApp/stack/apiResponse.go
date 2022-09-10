package stack

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}) (*events.APIGatewayV2HTTPResponse, error) {
	resp := events.APIGatewayV2HTTPResponse{Headers: map[string]string{"content-Type": "application/json"}}
	resp.StatusCode = status
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp.Body = string(jsonBody)
	return &resp, nil
}
