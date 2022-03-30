package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	var tests = []struct {
		rMethod  string
		usrAgent []string
		tName    string
		want     string
	}{
		{
			rMethod:  "GET",
			usrAgent: []string{"user-agent", "curl"},
			tName:    "test txt file",
			want:     "# Hallo\nthis is a simple txt file\n",
		},
		{
			rMethod:  "GET",
			usrAgent: []string{"user-agent", "mozilla"},
			tName:    "test html file",
			want:     "\n<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <title>42</title>\n</head>\n<body>\n\n<h1>Hello</h1>\n<p>This is a simple HTML template</p>\n\n</body>\n</html>\n\n",
		},
	}

	// #1
	for _, tt := range tests {
		req := httptest.NewRequest(
			tt.rMethod,
			"/",
			bytes.NewBuffer(nil),
		)
		req.Header.Set(
			tt.usrAgent[0],
			tt.usrAgent[1],
		)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(TestHandler)
		handler.ServeHTTP(rr, req)
		res := rr.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		got := string(data)
		want := tt.want

		t.Run(tt.tName, func(t *testing.T) {
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
