package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
	person1 := Person{Name: "John", Age: 30}
    jsonBytes, err := json.Marshal(person1)
    if err != nil {
        panic(err)
    }
    jsonString := string(jsonBytes)
    fmt.Println(jsonString) // {"name":"John","age":30}

    var person2 Person
    err = json.Unmarshal([]byte(jsonString), &person2) // you can also pass jsonBytes instead of []byte(jsonString)
    if err != nil {
        panic(err)
    }
    fmt.Println(person2.Name, person2.Age) // John 30
}
