package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	user := map[string]interface{}{
		"name": "Obj Go",
		"age": "Age do user",
		"date": "27/01/26",
	}

	userKeys := obj_json_func(user)

	fmt.Println(userKeys)

	userstring := `{"name": "Cassio", "date":"27/01/26"}`

	userParse := parseFunc(userstring)

	fmt.Println(userParse)

}

func obj_json_func(data map[string]interface{}) []string {
	keys := []string{}

	for k := range data {
		keys = append(keys, k)
	}

	return keys
}


//json.Unmarshal([]byte(jsonString), &obj) // decode e json.Marshal(obj) // encode

//obs json.Unmarshal retorna error por isso é importante apontar &

func parseFunc(data string) map[string]interface{} {
	user := make(map[string]interface{})

	json.Unmarshal([]byte(data), &user)

	return user
}
