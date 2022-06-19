package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func objectAndByteConversion() {
	payload := generatePayload()

	b := convertObjectToBytes(payload)
	convertedBackPayload := convertBytesToObject(b)

	fmt.Printf("%+v\n%+v\n", payload, convertedBackPayload)
}

func convertObjectToBytes(payload Payload) []byte {
	output, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

func convertBytesToObject(b []byte) Payload {
	payload := &Payload{}
	err := json.Unmarshal(b, payload)
	if err != nil {
		log.Fatal(err)
	}

	return *payload
}
