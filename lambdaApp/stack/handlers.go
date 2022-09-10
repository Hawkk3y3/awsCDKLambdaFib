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
	strIndex := req.QueryStringParameters["index"]
	if strIndex == "" {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String("invalid query param")})
	}
	intIndex, err := strconv.Atoi(strIndex)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	memo := make(map[int]int)
	result := pkg.CalculateFibUsingMemo(intIndex, memo)
	return apiResponse(200, ResultBody{aws.Int(result)})
}
