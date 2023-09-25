package httpclient

import (
	"bytes"
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/feelgood-inc/flgd-gommon/utils"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestFeelgoodResponseToStructOK(t *testing.T) {
	httpResponse := http.Response{
		Status:           "200 OK",
		StatusCode:       200,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           nil,
		Body:             nil,
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Uncompressed:     false,
		Trailer:          nil,
		Request:          nil,
		TLS:              nil,
	}
	response := resty.Response{
		Request:     nil,
		RawResponse: &httpResponse,
	}

	var appointment models.Appointment
	_ = utils.FeelgoodResponseToStruct(&response, &appointment)

	assert.IsType(t, models.Appointment{}, appointment)
}

func TestToRawFeelgoodResponseOK(t *testing.T) {
	body := `{"code": 200,"message": "OK","error": "","data": {"appointment": {"id": "5f9b1b1b-5b1a-4b0a-8b0a-5f9b1b1b5b1a"}}}`
	test := io.NopCloser(bytes.NewBufferString(body))

	response := &resty.Response{
		RawResponse: &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Header:     nil,
			Body:       test,
		},
	}

	raw, err := utils.ToRawFeelgoodResponse(response)

	assert.NoError(t, err)
	assert.Equal(t, 200, raw.Code)
}

func TestToRawFeelgoodResponseError(t *testing.T) {
	body := `{"code": 200,"message": "OK","error": "","data": {"appointment": {"id": "5f9b1b1b-5b1a-4b0a-8b0a-5f9b1b1b5b1a"}}}`

	response := &resty.Response{
		RawResponse: &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Header:     nil,
			Body:       io.NopCloser(bytes.NewBufferString(body)),
		},
	}

	_, err := utils.ToRawFeelgoodResponse(response)

	assert.Error(t, err)
}
