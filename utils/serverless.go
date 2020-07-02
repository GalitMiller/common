/*
 * Author: Galit Miller
 */

package utils

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// HTTPOK returns 200 in a format that the AWS API Gateway can understand
func HTTPOK(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

// HTTPServerError returns 500 in a format that the AWS API Gateway can understand
func HTTPServerError(err error) (events.APIGatewayProxyResponse, error) {
	Error("%s\n", err)
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// HTTPClientError returns 4xx in a format that the AWS API Gateway can understand
func HTTPClientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}
