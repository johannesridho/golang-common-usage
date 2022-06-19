package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func sendGetRequest() {
	resp, err := http.Get("https://reqres.in/api/users")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// encode result as map
	headers := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&headers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode, headers)
}

func sendPostRequest() {
	payload := generatePayload()

	// encode the payload
	buf := &bytes.Buffer{}
	err := json.NewDecoder(buf).Decode(&payload)

	req, err := http.NewRequest(http.MethodPost, "https://reqres.in/api/users", buf)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// decode response body
	respBody := &Response{}
	err = json.NewDecoder(resp.Body).Decode(respBody)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", respBody)
}
