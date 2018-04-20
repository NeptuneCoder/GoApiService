package product

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"net/url"
)

func TestGetOldProductInfo(t *testing.T) {
	w := httptest.NewRecorder()
	r := &http.Request{}
	GetOldProductInfo(w, r)
}
func TestPaymentInfo(t *testing.T) {
	w := httptest.NewRecorder()
	r := &http.Request{}
	header := http.Header{}
	header.Add("language", "CN")
	r.Header = header
	data := url.Values{}
	data.Add("key", "this is test data")
	r.Form = data
	PaymentInfo(w, r)
}
