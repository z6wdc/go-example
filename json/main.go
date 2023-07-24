package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Id    int    `json:"id"`
	Name  string `json:"-"`
	Mail  string `json:"mail"`
	Phone string `json:"phone"`
}

func main() {
	b := []byte(`{"id":1,"name":"aaa","mail":"aaa@example.com"}`)
	var u user
	if err := json.Unmarshal(b, &u); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(u)

	v, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(v))
}
