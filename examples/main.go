package main

import (
	"fmt"
	"log"

	fda "github.com/kitchenchem/go-openFda"
)

func main() {
	x, err := fda.NewClient("")
	if err != nil {
		log.Fatal(err)

	}

	var nilinterface any // just for testing purposes

	req, err := x.NewRequest("device/510k.json?search=advisory_committee:cv&limit=1", nilinterface)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("request method: %s\n\n", req.Method)

	var foo map[string]any // TODO works to avoid errors; need an actual not nil struct slice of structs (nil pointers ok) to write to
	res, err := x.Do(req, &foo)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range foo {

		fmt.Println(k, "value is ", v)
	}
	fmt.Println(res)
}
