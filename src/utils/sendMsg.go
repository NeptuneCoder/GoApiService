package utils

import (
	"net/http"
	"encoding/json"
	"fmt"
	"bytes"
	"github.com/yanghai23/GoLib/aterr"
	"io"
)

func SendNotify(url string, message string) func() (*http.Request, error) {
	return func() (*http.Request, error) {
		url := url
		msg := make(map[string]interface{})

		msg["msgtype"] = "text"

		content := make(map[string]string)
		content["content"] = message //"d8zone生成微信预支付订单接口出问题了"
		msg["text"] = content

		at := make(map[string]interface{})
		at["atMobiles"] = "[]"
		at["isAtAll"] = false
		msg["at"] = &at

		data, err := json.Marshal(msg)
		if err != nil {
			return nil, err
		}
		fmt.Println("data = ", string(data))

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		req.Header.Set("Content-type", "application/json;charset=UTF-8")
		return req, nil
	}
}
func OkStatus(w http.ResponseWriter, qr string) {
	result := make(map[string]interface{})
	result["code"] = 200
	result["msg"] = "提交成功"
	result["data"] = qr

	data, err := json.Marshal(result)
	fmt.Println(string(data))
	aterr.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(data))
}