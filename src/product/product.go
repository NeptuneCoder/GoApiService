package product

import (
	"net/http"
	"fmt"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

}

func PaymentInfo(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Header.Get("language"))
	fmt.Println("value:",r.Form.Get("key"))
}