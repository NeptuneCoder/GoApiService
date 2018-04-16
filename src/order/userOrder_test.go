package order

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestGetMyOrder(t *testing.T) {
	w := httptest.NewRecorder()
	r := &http.Request{}
	GetMyOrder(w, r)
}
