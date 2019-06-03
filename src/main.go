package main


import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
//	"github.com/aws/aws-sdk-go/aws"
//	"github.com/aws/aws-sdk-go/aws/session"
//	"github.com/aws/aws-sdk-go/service/s3/s3manager"
//	"fmt"
//	"os"
	"errors"
)


func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if (request.HTTPMethod=="GET") {
		ApiResponse :=events.APIGatewayProxyResponse{Body:"It actually works", StatusCode: 200};
		return ApiResponse, nil;
	}

	err:=errors.New("Method not allowed!");
	ApiResponse:=events.APIGatewayProxyResponse{Body:"Method not allowed", StatusCode: 405};
	return ApiResponse, err;



}

func main() {
	lambda.Start(handleRequest)
}


//func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
//	fmt.Printf("Body size = %d.\n", len(request.Body))
//
//	fmt.Println("Headers:")
//	for key, value := range request.Headers {
//		fmt.Printf("    %s: %s\n", key, value)
//	}
//
//	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
//}
//
//func main() {
//	lambda.Start(handleRequest)
//}