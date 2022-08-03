package httpclient

import (
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
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
	_ = FeelgoodResponseToStruct(&response, &appointment)

	assert.IsType(t, models.Appointment{}, appointment)
}
