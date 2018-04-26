package order

import (
	"net/http"
	"status/statusCode"
	"status/statusMsg"
	"mysessoin"
	"utils"
)

func GetMyOrder(w http.ResponseWriter, r *http.Request) {
	_, err := mysessoin.CheckSession(r)
	if err != nil {
		utils.RStatus(w, statusCode.LOGIN_INVALID, statusMsg.LOGIN_INVALID, "")
		return
	}
}
