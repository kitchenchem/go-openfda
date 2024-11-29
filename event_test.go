package fda

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestEvent(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/device/event.json",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{
                "results": [
                    {"adverse_event_flag": "Y"}
                ],
                "meta": null
            }`)
		},
	)

	response, resp, err := client.Event.GetEvent(&EventOptions{})
	if err != nil {
		t.Logf("Error: %v", err)

		if resp != nil && resp.Body != nil {
			body, readErr := io.ReadAll(resp.Body)
			if readErr == nil {
				t.Logf("Response body: %s", string(body))
			}
		}

		t.Fatalf("Event.GetEvent error: %v", err)
	}

	if len(response.Results) == 0 {
		t.Fatal("No results returned")
	}

	want := EventResponse{
		Results: []*Event{
			{AdverseEventFlag: "Y"},
		},
		Meta: nil,
	}

	if len(response.Results) != len(want.Results) {
		t.Errorf("Unexpected number of results. Got %d, want %d",
			len(response.Results), len(want.Results))
	}

	if response.Results[0].AdverseEventFlag != want.Results[0].AdverseEventFlag {
		t.Errorf("Unexpected Evaluation ID. Got %s, want %s",
			response.Results[0].AdverseEventFlag, want.Results[0].AdverseEventFlag)
	}

	if !reflect.DeepEqual(want, response) {
		t.Errorf("Event.GetEvent returned %+v, want %+v", response, want)
	}
}
