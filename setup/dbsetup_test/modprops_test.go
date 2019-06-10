package dbsetup_test

import (
	"testing"

	"github.com/andersonlira/stockids/model"
	"github.com/andersonlira/stockids/setup/dbsetup"
)

func TestGetModelProps(t *testing.T) {
	fields := 6
	model := model.Family{}
	props := dbsetup.GetModelProps(&model)

	if len(props) != fields {
		t.Errorf("Props size should be %d but %d", fields, len(props))
	}

	propID := dbsetup.ModelProp{}
	propName := dbsetup.ModelProp{}

	for _, p := range props {
		if p.FieldName == "id" {
			propID = p
		}
		if p.FieldName == "name" {
			propName = p
		}

	}

	if propID.FieldName != "id" {
		t.Errorf("FieldName should be 'id', but %s", propID.FieldName)
	}

	if !propID.FieldIndex {
		t.Errorf("FielIndex should be true, but %v", propID.FieldIndex)
	}

	if propName.FieldIndex {
		t.Errorf("FielIndex should be false, but %v", propName.FieldIndex)
	}

	if propID.FieldKeyType != "HASH" {
		t.Errorf("FieldName should be 'HASH', but %s", propID.FieldKeyType)
	}

	if propName.FieldKeyType != "" {
		t.Errorf("FieldName should be '', but %s", propName.FieldKeyType)
	}

}
