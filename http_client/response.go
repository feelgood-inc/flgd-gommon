package httpclient

import (
	"github.com/feelgood-inc/flgd-gommon/utils"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

func FeelgoodResponseToStruct(response *resty.Response, obj interface{}) error {
	responseAsBytes := gjson.GetBytes(response.Body(), "data")
	err := utils.JsonToStruct(responseAsBytes.Raw, &obj)
	if err != nil {
		return nil
	}

	return nil
}
