package utils

import (
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/tidwall/gjson"
)

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

func FeelgoodResponseToStruct(response *resty.Response, obj interface{}) error {
	responseAsBytes := gjson.GetBytes(response.Body(), "data")
	err := JsonToStruct(responseAsBytes.Raw, &obj)
	if err != nil {
		return err
	}

	return nil
}
