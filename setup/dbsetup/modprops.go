package dbsetup

import (
	"reflect"
	"strings"
)

//ModelProp is the AWS properties of a field in a
// giving interface
type ModelProp struct {
	FieldName    string
	FieldIndex   bool
	FieldKeyType string
}

//GetModelProps return a list of ModelProp of giving interface
func GetModelProps(m interface{}) []ModelProp {
	props := []ModelProp{}

	val := reflect.ValueOf(m).Elem()

	for i := 0; i < val.NumField(); i++ {

		typeField := val.Type().Field(i)
		tag := typeField.Tag
		isIndex, keyType := getDynamoProperties(tag.Get("dynamo"))

		props = append(props, ModelProp{
			FieldName:    tag.Get("json"),
			FieldIndex:   isIndex,
			FieldKeyType: keyType,
		})
	}

	return props
}

func getDynamoProperties(tag string) (isIndex bool, keyType string) {
	arr := strings.SplitN(tag, ";", -1)
	for _, s := range arr {
		if s == "hash" || s == "range" {
			keyType = strings.ToUpper(s)
			isIndex = true
			return
		}
	}
	return
}
