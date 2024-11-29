package fda

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestRecall(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/device/recall.json",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{
                "results": [
                    {"cfres_id": "30043"}
                ],
                "meta": null
            }`)
		},
	)

	response, resp, err := client.Recall.GetRecall(&RecallOptions{})
	if err != nil {
		t.Logf("Error: %v", err)

		if resp != nil && resp.Body != nil {
			body, readErr := io.ReadAll(resp.Body)
			if readErr == nil {
				t.Logf("Response body: %s", string(body))
			}
		}

		t.Fatalf("Recall.GetRecall error: %v", err)
	}

	if len(response.Results) == 0 {
		t.Fatal("No results returned")
	}

	want := RecallResponse{
		Results: []*Recall{
			{CfresID: "30043"},
		},
		Meta: nil,
	}

	if len(response.Results) != len(want.Results) {
		t.Errorf("Unexpected number of results. Got %d, want %d",
			len(response.Results), len(want.Results))
	}

	if response.Results[0].CfresID != want.Results[0].CfresID {
		t.Errorf("Unexpected Evaluation ID. Got %s, want %s",
			response.Results[0].CfresID, want.Results[0].CfresID)
	}

	if !reflect.DeepEqual(want, response) {
		t.Errorf("Recall.GetRecall returned %+v, want %+v", response, want)
	}
}
