package product

import (
	"net/http"
	"fmt"
	"utils"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

}

func PaymentInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Header.Get("language"))
	fmt.Println("value:", r.Form.Get("key"))

	utils.OkStatus(w, 200,"支付成功", "{\"key\":\"\"}")
}
