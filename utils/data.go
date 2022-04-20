package utils

import "encoding/json"

func StructToJson(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func JsonToStruct(jsonStr string, obj interface{}) error {
	return json.Unmarshal([]byte(jsonStr), obj)
}
