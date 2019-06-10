package family_test

import (
	"encoding/json"
	"testing"

	lt "github.com/andersonlira/stockids/lambdas_test"
	"github.com/andersonlira/stockids/model"
	gli "github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

func TestCreate(t *testing.T) {
	payload := lt.GetPayload("POST")
	response, err := gli.Run(gli.Input{
		Port:    8001,
		Payload: payload,
	})

	if err != nil {
		t.Errorf("Error was not expected here, but %v", err)
	}

	result := make(map[string]interface{})
	json.Unmarshal(response, &result)

	if v, ok := result["statusCode"]; !ok || v.(float64) != 400 {
		t.Errorf("Bad request expected, but %v", string(response))
	}

	family := model.Family{
		Name:        "Family Test",
		Description: "Family from unit test",
		Avatar:      "source/avatar",
	}
	body, _ := json.Marshal(family)
	payload["body"] = string(body)

	response, err = gli.Run(gli.Input{
		Port:    8001,
		Payload: payload,
	})

	if err != nil {
		t.Errorf("Error was not expected here, but %v", err)
	}

	result = make(map[string]interface{})
	json.Unmarshal(response, &result)

	if v, ok := result["statusCode"]; !ok || v.(float64) != 201 {
		t.Errorf("Bad request expected, but %v", string(response))
	}

}
