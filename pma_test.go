package fda

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestPma(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/device/pma.json",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{
                "results": [
                    {"decision_code": "OK30"}
                ],
                "meta": null
            }`)
		},
	)

	response, resp, err := client.Pma.GetPma(&PmaOptions{})
	if err != nil {
		t.Logf("Error: %v", err)

		if resp != nil && resp.Body != nil {
			body, readErr := io.ReadAll(resp.Body)
			if readErr == nil {
				t.Logf("Response body: %s", string(body))
			}
		}

		t.Fatalf("Pma.GetPma error: %v", err)
	}

	if len(response.Results) == 0 {
		t.Fatal("No results returned")
	}

	want := PmaResponse{
		Results: []*Pma{
			{DecisionCode: "OK30"},
		},
		Meta: nil,
	}

	if len(response.Results) != len(want.Results) {
		t.Errorf("Unexpected number of results. Got %d, want %d",
			len(response.Results), len(want.Results))
	}

	if response.Results[0].DecisionCode != want.Results[0].DecisionCode {
		t.Errorf("Unexpected Evaluation ID. Got %s, want %s",
			response.Results[0].DecisionCode, want.Results[0].DecisionCode)
	}

	if !reflect.DeepEqual(want, response) {
		t.Errorf("Pma.GetPma returned %+v, want %+v", response, want)
	}
}
