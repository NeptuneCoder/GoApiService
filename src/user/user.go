package user

import (
	"net/http"
	"status/statusCode"
	"status/statusMsg"
	"mysessoin"
	"utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	_, err := mysessoin.CheckSession(r)
	if err != nil {
		utils.OkStatus(w, statusCode.LOGIN_INVALID, statusMsg.LOGIN_INVALID, "")
		return
	}

}

func LoginOut(w http.ResponseWriter, r *http.Request) {
	_, err := mysessoin.CheckSession(r)
	if err != nil {
		utils.OkStatus(w, statusCode.LOGIN_INVALID, statusMsg.LOGIN_INVALID, "")
		return
	}
}
