package lambdas

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/mitchellh/mapstructure"
)

//Claims security model
type Claims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

//GetClaims return claims object form giving request
func GetClaims(request events.APIGatewayProxyRequest) (*Claims, error) {

	input := request.RequestContext.Authorizer["claims"]
	output := Claims{}
	err := mapstructure.Decode(input, &output)

	if err != nil {
		return nil, err
	}

	return &output, nil
}
