package fda

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestCovid19Serology(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/device/covid19serology.json",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{
                "results": [
                    {"evaluation_id": "K780531maf3378-a001"}
                ],
                "meta": null
            }`)
		},
	)

	response, resp, err := client.Covid19Serology.GetCovid19Serology(&Covid19SerologyOptions{})
	if err != nil {
		t.Logf("Error: %v", err)

		if resp != nil && resp.Body != nil {
			body, readErr := io.ReadAll(resp.Body)
			if readErr == nil {
				t.Logf("Response body: %s", string(body))
			}
		}

		t.Fatalf("Covid19Serology.GetCovid19Serology error: %v", err)
	}

	if len(response.Results) == 0 {
		t.Fatal("No results returned")
	}

	want := Covid19SerologyResponse{
		Results: []*Covid19Serology{
			{EvaluationID: "K780531maf3378-a001"},
		},
		Meta: nil,
	}

	if len(response.Results) != len(want.Results) {
		t.Errorf("Unexpected number of results. Got %d, want %d",
			len(response.Results), len(want.Results))
	}

	if response.Results[0].EvaluationID != want.Results[0].EvaluationID {
		t.Errorf("Unexpected Evaluation ID. Got %s, want %s",
			response.Results[0].EvaluationID, want.Results[0].EvaluationID)
	}

	if !reflect.DeepEqual(want, response) {
		t.Errorf("Covid19Serology.GetCovid19Serology returned %+v, want %+v", response, want)
	}
}
