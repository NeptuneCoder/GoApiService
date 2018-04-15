package product

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestGetProduct(t *testing.T) {
	w := httptest.NewRecorder()
	r := &http.Request{}
	GetProduct(w, r)

}

func TestPaymentInfo(t *testing.T) {
	w := httptest.NewRecorder()
	r := &http.Request{}
	PaymentInfo(w, r)
}
