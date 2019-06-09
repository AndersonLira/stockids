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

	// _, err := GetClaims(request)

	// if err == nil {
	// 	b := request.RequestContext
	// 	response, _ := json.Marshal(&b)
	// 	return addHeaders(events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusAccepted}, nil)
	// }

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
