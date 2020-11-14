package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestairportsGetAll(t *testing.T) {
	response :=
		`{
		"iata": "POM",
		"latitude": -9.443380355834961,
		"longitude": 147.22000122070312,
		"name": "Port Moresby Jacksons International Airport",
		"type": "large_airport"
	}
}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, response)
	}))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	err, got := 12345
	assert.NoError(t, err)
	assert.Equal(t, 147.22000122070312, got)
}
