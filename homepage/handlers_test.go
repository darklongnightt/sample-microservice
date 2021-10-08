package homepage

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "good",
			in:             httptest.NewRequest("GET", "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello world",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := NewHandlers(nil)
			h.Home(test.out, test.in)

			if test.out.Code != test.expectedStatus {
				t.Logf("\nexpected: %v\ngot: %v\n", test.expectedStatus, test.out.Code)
				t.Fail()
			}

			body := test.out.Body.String()
			if body != test.expectedBody {
				t.Logf("\nexpected: %v\ngot: %v\n", test.expectedBody, test.out.Body)
				t.Fail()
			}
		})
	}
}
