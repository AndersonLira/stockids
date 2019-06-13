package lambdas_test

import (
	"github.com/andersonlira/godyn/db"
	"github.com/andersonlira/godyn/model"
)

//TestHelper is the framework that tests must be executed
type TestHelper struct {
	//Tables to be created
	Tables []model.Entity
}

func (th *TestHelper) Setup() {
	for i := range th.Tables {
		db.CreateTable(th.Tables[i])
	}
}

func (th TestHelper) Teardown() {
	for i := range th.Tables {
		db.DeleteTable(th.Tables[i])
	}
}

//GetPayload returns default payloads for lambdas test
func GetPayload(method string) map[string]interface{} {
	payload := make(map[string]interface{})
	payload["httpMethod"] = method
	requestContext := make(map[string]interface{})
	authorizer := map[string]interface{}{
		"claims": map[string]string{
			"at_hash":          "pbrS5bw2kHgvKsUIs03Idw",
			"aud":              "2rsmg3ia773divrdaf8bjhcic4",
			"auth_time":        "1560184659",
			"cognito:groups":   "us-east-1_rhl2NnPzq_Facebook",
			"cognito:username": "Facebook_3186799091337591",
			"email":            "anderson4281@yahoo.com.br",
			"exp":              "Mon Jun 10 17:37:39 UTC 2019",
			"iat":              "Mon Jun 10 16:37:39 UTC 2019",
			"identities":       "{\"dateCreated\":\"1555241917584\",\"userId\":\"3186799091337591\",\"providerName\":\"Facebook\",\"providerType\":\"Facebook\",\"issuer\":null,\"primary\":\"true\"}",
			"iss":              "https://cognito-idp.us-east-1.amazonaws.com/us-east-1_rhl2NnPzq",
			"nonce":            "lmfFNQk-cNDR63-X-KWrjiccfEeUPwpN5O2RdA2fMPl7xXpdft0B8_vtXPOYqtv2n1fHsuHnslpKdVlmZkF74CPPuczV-tj9sg5q92_jSAZHdOiLXQ9M4Ozb4VX3RVGZ1KW_1Fb3GUVw71NoPI_kJjQothHqXE7x9qdUK8aoONM",
			"sub":              "e8a133db-017b-46fb-b6ad-04332e4b4b37",
			"token_use":        "id",
		},
	}
	requestContext["authorizer"] = authorizer
	payload["requestContext"] = requestContext
	return payload
}
