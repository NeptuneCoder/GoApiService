package user

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestLoginOut(t *testing.T) {
	r := &http.Request{}
	w := httptest.NewRecorder()
	LoginOut(w, r)
}
