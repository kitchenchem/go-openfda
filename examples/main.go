package main

import (
	"fmt"
	"log"
	"os"

	fda "github.com/kitchenchem/go-openFda"
)

func main() {
	x, err := fda.NewClient(os.Getenv("FDAKEY"))
	if err != nil {
		log.Fatal(err)

	}
	// path := fda.DevicePath + fda.F510k

	opt := &fda.Fda510kOptions{
		QueryParameters: fda.QueryParameters{
			Limit: "1",
		},
		AdvisoryCommittee: fda.Ptr("cv"),
	}

	f, _, err := x.Fda510k.Get510k(opt)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range f.Results {

		fmt.Printf("FDA510k: %+v\n", item)
	}

	fmt.Printf("Meta: %v\n", f.Meta.Results)
}
