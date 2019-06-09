package lambdas

import "github.com/aws/aws-lambda-go/events"

//Handlerable must implement basic methods for crud operation
type Handlerable interface {
	Get(request events.APIGatewayProxyRequest, claims Claims) (events.APIGatewayProxyResponse, error)
	Create(request events.APIGatewayProxyRequest, claims Claims) (events.APIGatewayProxyResponse, error)
	Update(request events.APIGatewayProxyRequest, claims Claims) (events.APIGatewayProxyResponse, error)
	Delete(request events.APIGatewayProxyRequest, claims Claims) (events.APIGatewayProxyResponse, error)
}
