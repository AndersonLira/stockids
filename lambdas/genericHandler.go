package lambdas

import "github.com/aws/aws-lambda-go/events"

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
	if request.HTTPMethod == "GET" {
		return gh.Handlerable.Get(request)
	}
	if request.HTTPMethod == "POST" {
		return gh.Handlerable.Create(request)
	}
	if request.HTTPMethod == "PUT" {
		return gh.Handlerable.Update(request)
	}
	if request.HTTPMethod == "DELETE" {
		return gh.Handlerable.Delete(request)
	}
	panic("GenericHandler is prepared to GET, POST, PUT and DELETE")
}
