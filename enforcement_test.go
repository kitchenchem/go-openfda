package fda

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestEnforcement(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/device/enforcement.json",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{
                "results": [
                    {"recall_number": "Z-1409-2013"}
                ],
                "meta": null
            }`)
		},
	)

	response, resp, err := client.Enforcement.GetEnforcement(&EnforcementOptions{})
	if err != nil {
		t.Logf("Error: %v", err)

		if resp != nil && resp.Body != nil {
			body, readErr := io.ReadAll(resp.Body)
			if readErr == nil {
				t.Logf("Response body: %s", string(body))
			}
		}

		t.Fatalf("Enforcement.GetEnforcement error: %v", err)
	}

	if len(response.Results) == 0 {
		t.Fatal("No results returned")
	}

	want := EnforcementResponse{
		Results: []*Enforcement{
			{RecallNumber: "Z-1409-2013"},
		},
		Meta: nil,
	}

	if len(response.Results) != len(want.Results) {
		t.Errorf("Unexpected number of results. Got %d, want %d",
			len(response.Results), len(want.Results))
	}

	if response.Results[0].RecallNumber != want.Results[0].RecallNumber {
		t.Errorf("Unexpected Evaluation ID. Got %s, want %s",
			response.Results[0].RecallNumber, want.Results[0].RecallNumber)
	}

	if !reflect.DeepEqual(want, response) {
		t.Errorf("Enforcement.GetEnforcement returned %+v, want %+v", response, want)
	}
}
