package utils

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Unsuport() (string, error) {
	var jsonInstance map[string]interface{} = make(map[string]interface{})
	jsonInstance["code"] = 200
	jsonInstance["message"] = "unsupport request"
	json, err := json.Marshal(jsonInstance)
	if err != nil {
		fmt.Println("json err: ", err)
		return "", errors.New("json error")
	}
	return string(json), nil
}