package lambdas

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

var headers = map[string]string{"Access-Control-Allow-Origin": "*"}

//GenericHandler is basic implementation for crud calls
type GenericHandler struct {
	Handlerable Handlerable
}

//Handler for start lambda
func (gh *GenericHandler) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if gh.Handlerable == nil {
		panic("Handlerable nil is not allowed")
	}

	c, err := GetClaims(request)
	if err != nil {
		return addHeaders(events.APIGatewayProxyResponse{Body: string("{\"message\":\"forbidden\""), StatusCode: http.StatusForbidden}, nil)
	}

	if c.Email == "" {
		panic("XXXXXsdfasdfa")
		return addHeaders(events.APIGatewayProxyResponse{Body: string("{\"message\":\"forbidden\""), StatusCode: http.StatusForbidden}, nil)
	}

	if request.HTTPMethod == "GET" {
		return addHeaders(gh.Handlerable.Get(request, *c))
	}
	if request.HTTPMethod == "POST" {
		return addHeaders(gh.Handlerable.Create(request, *c))
	}
	if request.HTTPMethod == "PUT" {
		return addHeaders(gh.Handlerable.Update(request, *c))
	}
	if request.HTTPMethod == "DELETE" {
		return addHeaders(gh.Handlerable.Delete(request, *c))
	}
	panic("GenericHandler is prepared to GET, POST, PUT and DELETE")
}

func addHeaders(response events.APIGatewayProxyResponse, err error) (events.APIGatewayProxyResponse, error) {
	response.Headers = headers
	return response, err
}
