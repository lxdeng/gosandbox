package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	/*
		Name string `json:"name"`
		Age  int    `json:"age"`
	*/

	Name string
	Age  int
}

func main() {
	fromJsonToStruct()

	fromStructToJson()

	fromJsonToMap()

	fromJsonToAny()
}

func fromJsonToStruct() {
	fmt.Println("fromJsonToStruct:")
	data := []byte(`
        {
            "name": "jack",
            "age": 18
       }
    `)

	var p Person
	err := json.Unmarshal(data, &p)

	fmt.Printf("p type is %T\n", p)

	if err == nil {
		fmt.Printf("No error: %v\n", p)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()
}

func fromStructToJson() {
	fmt.Println("fromStructToJson")
	p := Person{"tom", 20}
	data, err := json.Marshal(p)

	if err == nil {
		fmt.Printf("No error: %v\n", data)
		s := string(data)
		fmt.Printf("No error: %v\n", s)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()
}

func fromJsonToMap() {
	fmt.Println("fromJsonToMap")
	data := []byte(`
        {
            "name": "jack",
            "age": 18
       }
    `)

	var parsed map[string]interface{}
	err := json.Unmarshal(data, &parsed)

	if err == nil {
		fmt.Printf("No error: %v\n", parsed)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()
}

func fromJsonToAny() {
	fmt.Println("fromJsonToAny")
	data := []byte(`
        {
            "name": "jack",
            "age": 18
       }
    `)

	var parsed interface{}
	err := json.Unmarshal(data, &parsed)

	if err == nil {
		fmt.Printf("No error: %v\n", parsed)
		fmt.Printf("No error: parsed type %T\n", parsed)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()
}
