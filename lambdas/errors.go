package lambdas

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

//InvalidPathParam when path param is not present
func InvalidPathParam() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: string("PATH PARAM MUST BE PRESENT"), StatusCode: http.StatusBadRequest}, nil
}

//Conflict generic error when some conflict exists
func Conflict() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: "CONFLICT. DATA IS INVALID", StatusCode: http.StatusConflict}, nil
}

//BadRequest when body data is invalid
func BadRequest() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: string("Bad Request"), StatusCode: http.StatusBadRequest}, nil
}
