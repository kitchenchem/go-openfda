package main

import (
	"fmt"
	"log"
	"net/http"

	fda "github.com/kitchenchem/go-openFda"
)

func main() {
	x, err := fda.NewClient("")
	if err != nil {
		log.Fatal(err)

	}
	path := fda.DevicePath + fda.F510k

	params := &fda.FDA510kOptions{
		QueryParameters: fda.QueryParameters{
			Limit: "1",
		},
		AdvisoryCommittee: fda.Ptr("an"),
	}

	req, err := x.NewRequest(http.MethodGet, path, params)
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
