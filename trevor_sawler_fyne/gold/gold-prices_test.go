package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

// our actual test
func TestGold_GetPrices(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})
	g := Gold{
		Prices: nil,
		Client: client,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}

	// since its now using our custom client and not actually going to the network, we
	// can make a more useful test as we know exactly what json we should be getting back
	if p.Price != 2000.01 {
		t.Error("wrong price returned:", p.Price)
	}
}
