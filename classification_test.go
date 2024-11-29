package fda

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetClassification(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/device/classification.json",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{
                "results": [{
				"medical_specialty": "OR",
				"device_class": "2",
				"regulation_number": "888.3050"
				}
                ],
                "meta": null
            }`)
		},
	)

	response, resp, err := client.Classification.GetClassification(&ClassificationOptions{})
	if err != nil {
		t.Logf("Error: %v", err)

		if resp != nil && resp.Body != nil {
			body, readErr := io.ReadAll(resp.Body)
			if readErr == nil {
				t.Logf("Response body: %s", string(body))
			}
		}

		t.Fatalf("Classification.GetClassification error: %v", err)
	}

	if len(response.Results) == 0 {
		t.Fatal("No results returned")
	}

	want := ClassificationResponse{
		Results: []*Classification{
			{
				MedicalSpecialty: "OR",
				DeviceClass:      "2",
				RegulationNumber: "888.3050",
			},
		},
		Meta: nil,
	}

	if response.Results[0].MedicalSpecialty != want.Results[0].MedicalSpecialty ||
		response.Results[0].DeviceClass != want.Results[0].DeviceClass ||
		response.Results[0].RegulationNumber != want.Results[0].RegulationNumber {
		t.Errorf("Unexpected result details. Got %+v, want %+v",
			response.Results[0], want.Results[0])
	}
}
