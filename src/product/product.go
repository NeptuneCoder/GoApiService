package product

import (
	"net/http"
	"fmt"
	"utils"
)

func GetOldProductInfo(w http.ResponseWriter, r *http.Request) {

}
func GetProductInfo(w http.ResponseWriter, r *http.Request) {

}

func GenerateOrder(w http.ResponseWriter, r *http.Request)  {

	utils.OkStatus(w, 200,"支付成功", "{\"key\":\"\"}")
}

func PaymentInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Header.Get("language"))
	fmt.Println("value:", r.Form.Get("key"))

	utils.OkStatus(w, 200,"支付成功", "{\"key\":\"\"}")
}
