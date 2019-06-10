package family_test

import (
	"testing"

	lt "github.com/andersonlira/stockids/lambdas_test"
	gli "github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

var origin = `
"authorizer": {
	"claims": {
		"at_hash": "pbrS5bw2kHgvKsUIs03Idw",
		"aud": "2rsmg3ia773divrdaf8bjhcic4",
		"auth_time": "1560184659",
		"cognito:groups": "us-east-1_rhl2NnPzq_Facebook",
		"cognito:username": "Facebook_3186799091337591",
		"email": "anderson4281@yahoo.com.br",
		"exp": "Mon Jun 10 17:37:39 UTC 2019",
		"iat": "Mon Jun 10 16:37:39 UTC 2019",
		"identities": "{\"dateCreated\":\"1555241917584\",\"userId\":\"3186799091337591\",\"providerName\":\"Facebook\",\"providerType\":\"Facebook\",\"issuer\":null,\"primary\":\"true\"}",
		"iss": "https://cognito-idp.us-east-1.amazonaws.com/us-east-1_rhl2NnPzq",
		"nonce": "lmfFNQk-cNDR63-X-KWrjiccfEeUPwpN5O2RdA2fMPl7xXpdft0B8_vtXPOYqtv2n1fHsuHnslpKdVlmZkF74CPPuczV-tj9sg5q92_jSAZHdOiLXQ9M4Ozb4VX3RVGZ1KW_1Fb3GUVw71NoPI_kJjQothHqXE7x9qdUK8aoONM",
		"sub": "e8a133db-017b-46fb-b6ad-04332e4b4b37",
		"token_use": "id"
	}
}
`

func TestFamilyHandler(t *testing.T) {
	payload := lt.GetPayload()
	response, err := gli.Run(gli.Input{
		Port:    8001,
		Payload: payload,
	})

	if err != nil {
		t.Errorf("Error was not expected here, but %v", err)
	}

	if response != nil {
		t.Errorf("response: %v", string(response))
	}

}
