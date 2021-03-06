package lambdas

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

var headers = map[string]string{"Access-Control-Allow-Origin": "*"}

//Handlerable must implement basic methods for crud operation
type Handlerable interface {
	Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Update(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Delete(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

//GenericHandler is basic implementation for crud calls
type GenericHandler struct {
	Handlerable Handlerable
}

//Handler for start lambda
func (gh *GenericHandler) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if gh.Handlerable == nil {
		panic("Handlerable nil is not allowed")
	}
	token, ok := request.QueryStringParameters["token"]
	if !ok || token != "ianianso290801" {
		return addHeaders(events.APIGatewayProxyResponse{Body: string("{\"message\":\"forbidden\""), StatusCode: http.StatusForbidden}, nil)
	}
	if request.HTTPMethod == "GET" {
		return addHeaders(gh.Handlerable.Get(request))
	}
	if request.HTTPMethod == "POST" {
		return addHeaders(gh.Handlerable.Create(request))
	}
	if request.HTTPMethod == "PUT" {
		return addHeaders(gh.Handlerable.Update(request))
	}
	if request.HTTPMethod == "DELETE" {
		return addHeaders(gh.Handlerable.Delete(request))
	}
	panic("GenericHandler is prepared to GET, POST, PUT and DELETE")
}

func addHeaders(response events.APIGatewayProxyResponse, err error) (events.APIGatewayProxyResponse, error) {
	response.Headers = headers
	return response, err
}
