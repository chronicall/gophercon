package routing_test

import (
    "log"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/chronicall/gophercon/pkg/routing"
)


func TestBaseRouter(t *testing.T) {
    handler := routing.BaseRouter()

    ts := httptest.NewServer(handler)
    defer ts.Close()

    res, err := http.Get(ts.URL + "/home")
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode != http.StatusOK {
        t.Errorf("Wrong status code %d (%d expected)", res.StatusCode, http.StatusOK)
    }
}
