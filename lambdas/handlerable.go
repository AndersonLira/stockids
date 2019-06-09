package lambdas

import "github.com/aws/aws-lambda-go/events"

//Handlerable must implement basic methods for crud operation
type Handlerable interface {
	Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Update(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Delete(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}
