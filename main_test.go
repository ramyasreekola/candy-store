package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)
var candyDataURL = "https://candystore.zimpler.net/api/candy"
const mockCandyData = `
[
    {
        "name": "Annika",
        "candy": "Geisha",
        "eaten": 100,
        "date": "2023-12-06"
    },
    {
        "name": "Jonas",
        "candy": "Geisha",
        "eaten": 200,
        "date": "2023-08-15"
    },
    {
        "name": "Jonas",
        "candy": "Kexchoklad",
        "eaten": 100,
        "date": "2023-05-28"
    },
    {
        "name": "Aadya",
        "candy": "Nötchoklad",
        "eaten": 2,
        "date": "2023-10-19"
    }
]
`

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Time
		wantErr bool
	}{
		{
			name:    "valid date",
			input:   `"2023-01-02"`,
			want:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "invalid date",
			input:   `"not-a-date"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ct CustomTime
			err := ct.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !ct.Time.Equal(tt.want) {
				t.Errorf("CustomTime.UnmarshalJSON() got = %v, want %v", ct.Time, tt.want)
			}
		})
	}
}
func Test_fetchCandyData(t *testing.T) {
    // Start a local HTTP server to mock the external API
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")
        io.WriteString(w, mockCandyData) 
    }))
    defer server.Close()

    oldURL := candyDataURL 
    candyDataURL = server.URL
    defer func() { candyDataURL = oldURL }()

    data, err := fetchCandyData()
    if err != nil {
        t.Errorf("fetchCandyData() error = %v, wantErr nil", err)
    }

    var candyData CandyDataSlice
    if err := json.Unmarshal(data, &candyData); err != nil {
        t.Fatalf("Error unmarshaling response data: %v", err)
    }

	if len(candyData) == 0 {
        t.Errorf("Expected non-empty candy data")
    }
}
