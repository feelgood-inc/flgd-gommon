package utils

import (
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/tidwall/gjson"
)

type RawFeelgoodResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

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

func ToRawFeelgoodResponse(response *resty.Response) (RawFeelgoodResponse, error) {
	var rawResponse RawFeelgoodResponse
	err := json.Unmarshal(response.Body(), &rawResponse)
	if err != nil {
		return RawFeelgoodResponse{}, err
	}

	return rawResponse, nil
}

func FeelgoodResponseToStruct(response *resty.Response, obj interface{}) error {
	responseAsBytes := gjson.GetBytes(response.Body(), "data")
	err := JsonToStruct(responseAsBytes.Raw, &obj)
	if err != nil {
		return err
	}

	return nil
}

func (r *RawFeelgoodResponse) IsZero() bool {
	return r == (&RawFeelgoodResponse{})
}

func (r *RawFeelgoodResponse) IsError() bool {
	return r.Error != ""
}
