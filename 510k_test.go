package fda

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGet510k(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/device/510k.json",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{
                "results": [
                    {"k_number": "K780531"}
                ],
                "meta": null
            }`)
		},
	)

	response, resp, err := client.Fda510k.Get510k(&Fda510kOptions{})
	if err != nil {
		t.Logf("Error: %v", err)

		if resp != nil && resp.Body != nil {
			body, readErr := io.ReadAll(resp.Body)
			if readErr == nil {
				t.Logf("Response body: %s", string(body))
			}
		}

		t.Fatalf("Fda510k.Get510k error: %v", err)
	}

	if len(response.Results) == 0 {
		t.Fatal("No results returned")
	}

	want := Fda510kResponse{
		Results: []*Fda510k{
			{KNumber: "K780531"},
		},
		Meta: nil,
	}

	if len(response.Results) != len(want.Results) {
		t.Errorf("Unexpected number of results. Got %d, want %d",
			len(response.Results), len(want.Results))
	}

	if response.Results[0].KNumber != want.Results[0].KNumber {
		t.Errorf("Unexpected K number. Got %s, want %s",
			response.Results[0].KNumber, want.Results[0].KNumber)
	}
}
