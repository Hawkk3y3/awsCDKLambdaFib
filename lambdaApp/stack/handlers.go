package stack

import (
	"github.com/Hawkk3y3/lambdaApp/pkg"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"

	"net/http"
	"strconv"
)

type ErrorBody struct {
	ErrorMsg *string `json:"error"`
}
type ResultBody struct {
	Result *int `json:"result"`
}

func Handler(req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	return CalFib(req)
}

func CalFib(req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	strNumber := req.QueryStringParameters["number"]
	if strNumber == "" {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String("invalid query param")})
	}
	intNumber, err := strconv.Atoi(strNumber)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	memo := make(map[int]int)
	result := pkg.CalculateFibUsingMemo(intNumber, memo)
	return apiResponse(200, ResultBody{aws.Int(result)})
}
