/*
 * Author: Galit Miller
 *
 * Copyright (c) 2020 RemoteCourthouse as an unpublished work.
 * All rights reserved.
 *
 * The information contained herein is confidential property of RemoteCourthouse.
 * The use, copying, transfer or disclosure of such
 * information is prohibited except by express written agreement with RemoteCourthouse.
 */

package libs

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// HttpOK returns 200 in a format that the AWS API Gateway can understand
func HttpOK(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

// HttpServerError returns 500 in a format that the AWS API Gateway can understand
func HttpServerError(err error) (events.APIGatewayProxyResponse, error) {
	Error("%s\n", err)
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// HttpClientError returns 4xx in a format that the AWS API Gateway can understand
func HttpClientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}
