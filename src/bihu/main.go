package main

import (
	_ "initialize"
	"github.com/yanghai23/GoLib/atsafe"
	"net/http"
	"feedback"
	"event"
	"product"
	"user"
)

func main() {

	//意见反馈的接口
	http.HandleFunc("/logUpload", Safe(feedback.UploadLog))
	http.HandleFunc("/findLogFile", Safe(feedback.FindLogFile))
	http.HandleFunc("/getFileContent", Safe(feedback.GetFile))
	http.HandleFunc("/findProblem", Safe(feedback.FindProblemContent))
	http.HandleFunc("/find/content", Safe(feedback.LastNewContent))

	//事件反馈的接口
	http.HandleFunc("/payment/info", Safe(event.PaymentStatus))
	http.HandleFunc("/save/event", Safe(event.SaveEvent))

	//个人账号相关的接口
	http.HandleFunc("/user/oauthLogin", Safe(user.Login))

	//商品先关的接口
	http.HandleFunc("/product/defineService", Safe(product.DefineNewService))
	//新的点单接口
	http.HandleFunc("/serve/showSetMeal2", Safe(product.GetProductInfo))
	//老的商品接口
	http.HandleFunc("/order/getProductInfo", Safe(product.GetOldProductInfo))
	http.HandleFunc("/order/generateOrder", Safe(product.GenerateOrder))

	http.ListenAndServe(":3000", nil)
}

func Safe(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return atsafe.SafeHandle(f)
}
