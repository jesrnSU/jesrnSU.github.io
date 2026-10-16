package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// Exercise the real router, decoder, service, and response together.
// httptest does this in memory: no listening port or database is needed.
func TestCalculateRoundTotal(t *testing.T) {
	payload, err := os.ReadFile("../../examples/round.json")
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/api/v1/rounds/total", strings.NewReader(string(payload)))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	NewRouter().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", response.Code, response.Body.String())
	}
	var body struct {
		TotalStrokes int `json:"total_strokes"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body.TotalStrokes != 7 {
		t.Fatalf("total_strokes = %d, want 7", body.TotalStrokes)
	}
}

func TestCalculateParTotal(t *testing.T) {
	payload, err := os.ReadFile("../../examples/round.json")
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/api/v1/rounds/course_par", strings.NewReader(string(payload)))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	NewRouter().ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", response.Code, response.Body.String())
	}
	var body struct {
		TotalPar int `json:"total_par"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body.TotalPar != 6 {
		t.Fatalf("total_par = %d, want 6", body.TotalPar)
	}

}

func TestCalculateRoundTotalRejectsInvalidJSON(t *testing.T) {
	for _, payload := range []string{`{`, `{"holes":[{"strokes":"four"}]}`} {
		t.Run(payload, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/api/v1/rounds/total", strings.NewReader(payload))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			NewRouter().ServeHTTP(response, request)
			if response.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400", response.Code)
			}
		})
	}
}

func TestScoringEndpointsRejectInvalidRound(t *testing.T) {
	for _, path := range []string{"/api/v1/rounds/total", "/api/v1/rounds/course_par"} {
		t.Run(path, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, path, strings.NewReader(
				`{"course_name":"Practice","date":"2026-10-16","holes":[{"hole_number":1,"par":3,"strokes":-1}]}`))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			NewRouter().ServeHTTP(response, request)
			if response.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400", response.Code)
			}
			var body struct {
				Error string `json:"error"`
			}
			if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			if body.Error != "hole 1: strokes must be positive" {
				t.Fatalf("unexpected error: %q", body.Error)
			}
		})
	}
}
