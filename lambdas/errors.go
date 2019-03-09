package lambdas

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

//InvalidPathParam when path param is not present
func InvalidPathParam() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: string("PATH PARAM MUST BE PRESENT"), StatusCode: http.StatusBadRequest}, nil
}
