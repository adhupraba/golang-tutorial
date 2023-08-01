package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page int `json:"page"`
	// Words []string `json:"words,omitempty"`
	Words []string `json:"words"`
}

func main() {
	r := Response{
		Page: 1,
		// Words: []string{"up", "in", "out"},
	}
	fmt.Printf("%#v\n", r)

	// encode into json string from struct
	j, _ := json.Marshal(r)
	fmt.Println(string(j))

	var r2 Response

	// decode from json string into struct
	json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)
}
