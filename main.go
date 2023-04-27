package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				buf.WriteString(line)
				break // end of the input
			} else {
				fmt.Println(err.Error())
				os.Exit(1) // something bad happened
			}
		}
		buf.WriteString(line)

	}

	fmt.Printf("valid json? %v\n", json.Valid(buf.Bytes()))

	type MedicalRecord struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var record MedicalRecord
	err := json.Unmarshal(buf.Bytes(), &record)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) // something bad happened
	}

	fmt.Printf("name: %s, age: %d\n", record.Name, record.Age)
}
