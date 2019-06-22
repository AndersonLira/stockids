package child_test

import (
	"encoding/json"
	"strings"
	"testing"

	md "github.com/andersonlira/godyn/model"
	lt "github.com/andersonlira/stockids/lambdas_test"
	"github.com/andersonlira/stockids/model"
	gli "github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

var childID string
var helper = lt.TestHelper{Tables: []md.Entity{&model.Child{}}}

func TestCreateChild(t *testing.T) {
	helper.Setup()
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

	child := model.Child{
		Name:        "Child Test",
		Description: "Child from unit test",
		Avatar:      "source/avatar",
	}
	body, _ := json.Marshal(child)
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
		t.Errorf("Created response expected, but %v", string(response))
	}

}

func TestGetChild(t *testing.T) {
	payload := lt.GetPayload("GET")
	response, err := gli.Run(gli.Input{
		Port:    8001,
		Payload: payload,
	})

	if err != nil {
		t.Errorf("Error was not expected here, but %v", err)
	}

	result := make(map[string]interface{})
	json.Unmarshal(response, &result)

	if v, ok := result["statusCode"]; !ok || v.(float64) != 200 {
		t.Errorf("Bad request expected, but %v", string(response))
	}
	body := result["body"]
	list := []model.Child{}
	json.Unmarshal([]byte(body.(string)), &list)
	if len(list) != 1 {
		t.Errorf("List should be an element, but %v", list)
	}

	item := list[0]
	if item.Name != "Child Test" {
		t.Errorf("Child name should be 'Child Test' but %s", item.Name)
	}

	childID = item.ID

}

func TestUpdateChild(t *testing.T) {
	payload := lt.GetPayload("PUT")
	pathParameters := make(map[string]string)
	pathParameters["id"] = childID
	payload["pathParameters"] = pathParameters

	child := model.Child{
		Name:        "Child Test Altered",
		Description: "Child from unit test",
		Avatar:      "source/avatar",
	}
	body, _ := json.Marshal(child)
	payload["body"] = string(body)

	response, err := gli.Run(gli.Input{
		Port:    8001,
		Payload: payload,
	})

	if err != nil {
		t.Errorf("Error was not expected here, but %v", err)
	}

	result := make(map[string]interface{})
	json.Unmarshal(response, &result)

	v, ok := result["statusCode"]

	if !ok || v.(float64) != 200 {
		t.Errorf("Update response expected, but %v", string(response))
	}

	resultBody := (result["body"]).(string)

	if strings.Index(resultBody, "Child Test Altered") < 0 {
		t.Errorf("Body should have 'Child Test Altered' but %s", resultBody)
	}

}

func TestDeleteChild(t *testing.T) {
	defer helper.Teardown()
	payload := lt.GetPayload("DELETE")
	pathParameters := make(map[string]string)
	pathParameters["id"] = childID
	payload["pathParameters"] = pathParameters

	response, err := gli.Run(gli.Input{
		Port:    8001,
		Payload: payload,
	})

	if err != nil {
		t.Errorf("Error was not expected here, but %v", err)
	}

	result := make(map[string]interface{})
	json.Unmarshal(response, &result)

	v, ok := result["statusCode"]

	if !ok || v.(float64) != 200 {
		t.Errorf("Update response expected, but %v", string(response))
	}

	resultBody := (result["body"]).(string)

	if strings.Index(resultBody, "true") < 0 {
		t.Errorf("Body should have true value but %s", resultBody)
	}

}
