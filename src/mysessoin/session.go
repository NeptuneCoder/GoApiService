package mysessoin

import (
	"net/http"
	"errors"
)

var userStatus = make(map[string]Cookie)

type Cookie struct {
	AccountId string
	Token     string
	Time      string
	ValidTime string
}

func CheckSession(r *http.Request) error {
	r.ParseForm()
	//获取头信息，token
	token := r.Header.Get("authorization")
	//查询缓存中是否有该信息
	cookie := userStatus[token]
	if cookie.AccountId == "" {
		//查询数据库
	}
	//查询数据库后，还是为空，则反馈错误

	//判断是否登录状态有效

	//如果无效，返回异常
	return errors.New("该用户登录过期,请重新登录")

}
