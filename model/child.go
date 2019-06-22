package model

import (
	"os"
)

//Child represents one group
type Child struct {
	ID          string `json:"id" dynamo:"hash"`
	CreatedAt   int64  `json:"created_at"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	UserID      string `json:"user_id" dynamo:"range"`
}

//GetTableName returns table name
func (Child) GetTableName() string {
	return "skChild" + os.Getenv("SK_ENVIRONMENT")
}
