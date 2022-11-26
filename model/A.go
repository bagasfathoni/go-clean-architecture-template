package model

import (
	"encoding/json"
	"time"
)

/*
	Model is like an object which modeled from a table in our database.
	In Go we can also choose to migrate the model to the database
*/

// Define every field in the struct based on the column
type A struct {
	Id        int
	Name      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Create "table_a" in the database
func (A) TableName() string {
	return "table_a"
}

// Turn the model into JSON
func (a *A) ToJSON() string {
	res, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		return ""
	}
	return string(res)
}
