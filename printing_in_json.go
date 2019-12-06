package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Arc struct {
	Head     int
	Modifier int
}

func main() {
	//arc := Arc{"saw", "He"}
	arc := Arc{}
	//fmt.Printf("%v\n", arc)
	//fmt.Printf("%+v\n", arc)
	//fmt.Printf("%#v\n", arc)

	// Convert structs to JSON.
	//data, err := json.Marshal(arc)
  data, err := json.MarshalIndent(arc, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
