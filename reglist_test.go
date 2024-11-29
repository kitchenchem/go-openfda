package fda

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestReglist(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/device/reglist.json",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{
                "results": [
                    {"pma_number": ""}
                ],
                "meta": null
            }`)
		},
	)

	response, resp, err := client.Reglist.GetReglist(&ReglistOptions{})
	if err != nil {
		t.Logf("Error: %v", err)

		if resp != nil && resp.Body != nil {
			body, readErr := io.ReadAll(resp.Body)
			if readErr == nil {
				t.Logf("Response body: %s", string(body))
			}
		}

		t.Fatalf("Reglist.GetReglist error: %v", err)
	}

	if len(response.Results) == 0 {
		t.Fatal("No results returned")
	}

	want := ReglistResponse{
		Results: []*Reglist{
			{PMANumber: ""},
		},
		Meta: nil,
	}

	if len(response.Results) != len(want.Results) {
		t.Errorf("Unexpected number of results. Got %d, want %d",
			len(response.Results), len(want.Results))
	}

	if response.Results[0].PMANumber != want.Results[0].PMANumber {
		t.Errorf("Unexpected Evaluation ID. Got %s, want %s",
			response.Results[0].PMANumber, want.Results[0].PMANumber)
	}

	if !reflect.DeepEqual(want, response) {
		t.Errorf("Reglist.GetReglist returned %+v, want %+v", response, want)
	}
}
