package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

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

func readFile() {
	file, err := os.Open("files/file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
