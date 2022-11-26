package dto

import "encoding/json"

/*
	A Data Transfer Object is an object that is used to encapsulate data, and send it from one subsystem of an application to another.
*/

type AandB struct {
	Name   string
	Status bool
	Count  int
}

func (a *AandB) ToJSON() string {
	res, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		return ""
	}
	return string(res)
}
