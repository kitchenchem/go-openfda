package fda

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TesUdi(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/device/udi.json",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{
                "results": [
                    {"brand_name": "brand"}
                ],
                "meta": null
            }`)
		},
	)

	response, resp, err := client.Udi.GetUdi(&UdiOptions{})
	if err != nil {
		t.Logf("Error: %v", err)

		if resp != nil && resp.Body != nil {
			body, readErr := io.ReadAll(resp.Body)
			if readErr == nil {
				t.Logf("Response body: %s", string(body))
			}
		}

		t.Fatalf("Udi.GetUdi error: %v", err)
	}

	if len(response.Results) == 0 {
		t.Fatal("No results returned")
	}

	want := UdiResponse{
		Results: []*Udi{
			{BrandName: "brand"},
		},
		Meta: nil,
	}

	if len(response.Results) != len(want.Results) {
		t.Errorf("Unexpected number of results. Got %d, want %d",
			len(response.Results), len(want.Results))
	}

	if response.Results[0].BrandName != want.Results[0].BrandName {
		t.Errorf("Unexpected Evaluation ID. Got %s, want %s",
			response.Results[0].BrandName, want.Results[0].BrandName)
	}

	if !reflect.DeepEqual(want, response) {
		t.Errorf("Udi.GetUdi returned %+v, want %+v", response, want)
	}
}
