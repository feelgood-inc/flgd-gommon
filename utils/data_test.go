package utils

import (
	"bytes"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

type TestStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestStructToJson(t *testing.T) {
	// Test valid struct to JSON conversion
	obj := TestStruct{Name: "John Doe", Age: 30}
	expectedJSON := `{"name":"John Doe","age":30}`
	assert.Equal(t, expectedJSON, StructToJson(obj))

	// Test invalid struct to JSON conversion
	invalidObj := make(chan int) // Invalid type for JSON marshaling
	assert.Equal(t, "", StructToJson(invalidObj))
}

func TestJsonToStruct(t *testing.T) {
	// Test valid JSON to struct conversion
	jsonStr := `{"name":"Jane Doe","age":25}`
	expectedObj := TestStruct{Name: "Jane Doe", Age: 25}

	var obj TestStruct
	err := JsonToStruct(jsonStr, &obj)
	assert.Nil(t, err)
	assert.Equal(t, expectedObj, obj)

	// Test invalid JSON to struct conversion
	invalidJSON := `{"name":"Jack Doe","age":"invalid"}` // Invalid "age" type
	err = JsonToStruct(invalidJSON, &obj)
	assert.NotNil(t, err)
}

func TestFeelgoodResponseToStruct_ValidResponse(t *testing.T) {
	// Test valid JSON response to struct conversion
	jsonStr := `{"data": {"name":"Alice","age":40}}`
	expectedObj := TestStruct{Name: "Alice", Age: 40}

	response := &resty.Response{
		RawResponse: &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Header:     nil,
			Body:       ioutil.NopCloser(bytes.NewBufferString(jsonStr)),
		},
	}
	var obj TestStruct
	err := FeelgoodResponseToStruct(response, &obj)
	assert.Nil(t, err)
	assert.Equal(t, expectedObj, obj)
}

func TestFeelgoodResponseToStruct_InvalidResponse(t *testing.T) {
	// Test invalid JSON response to struct conversion
	invalidJSONStr := `{"data": {"name":"Bob","age":"invalid"}}` // Invalid "age" type
	invalidResponse := &resty.Response{
		RawResponse: &http.Response{
			Status:     "500 Internal Server Error",
			StatusCode: 500,
			Header:     nil,
			Body:       ioutil.NopCloser(bytes.NewBufferString(invalidJSONStr)),
		},
	}
	var obj TestStruct
	err := FeelgoodResponseToStruct(invalidResponse, &obj)
	assert.NotNil(t, err)
}
