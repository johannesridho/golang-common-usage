package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	stringBuilder()
	unicode()
}

func stringBuilder() {
	var strBuilder strings.Builder
	strBuilder.WriteString("add string ")
	strBuilder.WriteByte('b')
	strBuilder.WriteRune('c')

	fmt.Println(strBuilder.String())
}

func stringIntConversion() {
	str := "320"

	// convert to int
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	// convert to string
	str = strconv.Itoa(i)
}

func unicode() {
	str := "123"
	thirdChar := int(str[2] - '0')
	fmt.Println("third char", thirdChar, int(str[2]), int('0'))
}
