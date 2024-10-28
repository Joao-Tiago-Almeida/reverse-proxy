package utils

import "encoding/json"

func Map(data interface{}) map[string]interface{} {
	// Marshal the struct into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Unmarshal the JSON into a map
	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		panic(err)
	}

	return result
}

type Validatable interface {
	validate() bool
}

func ValidMap(data Validatable) map[string]interface{} {
	err := data.validate()
	if !err {
		panic(err)
	}

	return Map(data)
}
