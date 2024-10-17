package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	// "testing"
    // "task/findpairs/models"
	"net/http/httptest"
	"testing"
    "task/findpairs/models"
)

func TestFindPairs(t *testing.T) {
	tests := []struct {
		name       string
		input      models.Request
		expected   [][]int
		statusCode int
	}{
		{
			name: "Valid input with pairs",
			input: models.Request{
				Numbers: []int{1, 2, 3, 4, 5},
				Target:  6,
			},
			expected:   [][]int{{0, 4}, {1, 3}},
			statusCode: http.StatusOK,
		},
		{
			name: "Empty array",
			input: models.Request{
				Numbers: []int{},
				Target:  6,
			},
			expected:   nil,
			statusCode: http.StatusBadRequest,
		},
		{
			name: "No pairs found",
			input: models.Request{
				Numbers: []int{1, 2, 3},
				Target:  7,
			},
			expected:   [][]int{},
			statusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/find-pairs", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()

			FindPairs(rec, req)

			res := rec.Result()
			if res.StatusCode != tt.statusCode {
				t.Errorf("expected status %d, got %d", tt.statusCode, res.StatusCode)
			}

			var response map[string]interface{}
			json.NewDecoder(res.Body).Decode(&response)

			if res.StatusCode == http.StatusOK {
				if solutions, ok := response["solutions"].([]interface{}); ok {
					// Check if solutions match
					if len(solutions) != len(tt.expected) {
						t.Errorf("expected %v, got %v", tt.expected, solutions)
					}}}
		})
}}
