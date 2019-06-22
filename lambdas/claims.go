package lambdas

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/mitchellh/mapstructure"
)

//Claims security model
type Claims struct {
	Email    string `json:"email"`
	Username string `json:"cognito:username"`
}

//GetClaims return claims object form giving request
func GetClaims(request events.APIGatewayProxyRequest) (*Claims, error) {

	input := request.RequestContext.Authorizer["claims"].(map[string]interface{})
	output := Claims{}
	err := mapstructure.Decode(input, &output)

	//TODO how to get the property straight
	if v, ok := input["cognito:username"]; ok {
		output.Username = v.(string)
	}

	if err != nil {
		return nil, err
	}

	return &output, nil
}
