package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Payload struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

type Response struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

func main() {
	sendGetRequest()
	sendPostRequest()
	objectAndByteConversion()
	writeJSONToFile()
	writeToFile()
	writeToCSV()
}

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

func writeJSONToFile() {
	payloads := generatePayloads()

	file, err := os.Create("files/payloads.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(&payloads)
	if err != nil {
		log.Fatal(err)
	}
}

func writeToFile() {
	file, err := os.Create("files/file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wr := bufio.NewWriter(file)
	defer wr.Flush()

	for i := 0; i < 10; i++ {
		_, err = wr.WriteString("str" + strconv.Itoa(i) + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func writeToCSV() {
	payloads := generatePayloads()

	file, err := os.Create("files/payloads.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	err = wr.Write([]string{"name", "job"})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range payloads {
		err = wr.Write([]string{p.Name, p.Job})
		if err != nil {
			log.Fatal(err)
		}
	}
}
