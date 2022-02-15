package main

import (
	"encoding/json"
	"fmt"
)

type AutoGenerated struct {
	Age int `json:"age"`
	Name string `json:"name"`
	Child []int `json:"child"`
}

func main() {
	a := AutoGenerated{}

	jsonStr1 := `{"age": 14,"name": "potter", "child":[1,2,3]}`
	json.Unmarshal([]byte(jsonStr1),&a)
	aa := a.Child
	fmt.Println(aa)
	jsonStr2 := `{"age": 14,"name": "potter", "child":[3,4,5,7,8,9]}`
	json.Unmarshal([]byte(jsonStr2),&a)
	fmt.Println(aa)


}