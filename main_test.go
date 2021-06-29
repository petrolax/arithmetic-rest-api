package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestAPI(t *testing.T) {
	expectedres := []struct {
		request string
		result  map[string]interface{}
	}{
		{
			request: "http://0.0.0.0:8081/api/add?a=5&b=2",
			result: map[string]interface{}{
				"Success": true,
				"ErrCode": "",
				"Value":   float64(7),
			},
		},
		{
			request: "http://0.0.0.0:8081/api/sub?a=5&b=2",
			result: map[string]interface{}{
				"Success": true,
				"ErrCode": "",
				"Value":   float64(3),
			},
		},
		{
			request: "http://0.0.0.0:8081/api/mul?a=5&b=2",
			result: map[string]interface{}{
				"Success": true,
				"ErrCode": "",
				"Value":   float64(10),
			},
		},
		{
			request: "http://0.0.0.0:8081/api/div?a=5&b=2",
			result: map[string]interface{}{
				"Success": true,
				"ErrCode": "",
				"Value":   float64(2.5),
			},
		},
		{
			request: "http://0.0.0.0:8081/api/div?a=5&b=0",
			result: map[string]interface{}{
				"Success": false,
				"ErrCode": "You can't divide by zero",
				"Value":   nil,
			},
		},
		{
			request: "http://0.0.0.0:8081/api/add?a=5avs&b=0",
			result: map[string]interface{}{
				"Success": false,
				"ErrCode": "Parameter 'a' is not a number",
				"Value":   nil,
			},
		},
		{
			request: "http://0.0.0.0:8081/api/div?a=5&b=",
			result: map[string]interface{}{
				"Success": false,
				"ErrCode": "One of the parameters haven't value or didn't exist",
				"Value":   nil,
			},
		},
		{
			request: "http://0.0.0.0:8081/api/div?a=5",
			result: map[string]interface{}{
				"Success": false,
				"ErrCode": "One of the parameters haven't value or didn't exist",
				"Value":   nil,
			},
		},
	}
	for _, want := range expectedres {
		req, err := http.NewRequest("GET", want.request, nil)
		if err != nil {
			t.Errorf("Can't create new request: %s\n. Error Message: %s\n\n", want.request, err.Error())
			continue
		}

		req.Header.Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
		req.Header.Set("Access-Control-Allow-Header", "*")
		req.Header.Set("Access-Control-Allow-Credentials", "true")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("An error occurred while executing the request: %s\n Error Message: %s\n\n", want.request, err.Error())
			continue
		}

		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)
		// if res["Value"] != want.result["Value"] {
		if !reflect.DeepEqual(res, want.result) {
			t.Errorf("Wrong answer: %v", res)
		}
	}
}
