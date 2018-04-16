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

	http.HandleFunc("/logUpload", Safe(feedback.UploadLog))
	http.HandleFunc("/findLogFile", Safe(feedback.FindLogFile))
	http.HandleFunc("/getFileContent", Safe(feedback.GetFile))
	http.HandleFunc("/findProblem", Safe(feedback.FindProblemContent))
	http.HandleFunc("/find/content", Safe(feedback.LastNewContent))

	http.HandleFunc("/payment/info", Safe(event.PaymentStatus))
	http.HandleFunc("/save/event", Safe(event.SaveEvent))

	http.HandleFunc("/serve/showSetMeal2", Safe(product.GetProduct))

	//个人账号相关的接口
	http.HandleFunc("/user/oauthLogin", Safe(user.Login))

	http.ListenAndServe(":3000", nil)
}

func Safe(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return atsafe.SafeHandle(f)
}
