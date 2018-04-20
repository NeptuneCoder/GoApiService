package mysessoin

import (
	"net/http"
	"status/statusMsg"
	"fmt"
	"unsafe"
)

var userStatus = make(map[string]Cookie)

type Cookie struct {
	AccountId string
	Token     string
	Time      string
	ValidTime string
}

type Res struct {
	ErrMsg  string
	ErrCode int
}

//TODO 需要处理不同的请求类型过来时，获取token值
func CheckSession(r *http.Request) (res *Res, err error) {
	if r.Method == "POST" {
		return normal(r, func() {
			r.ParseForm()
		})
	}
	return nil, nil
}

//函数作为参数只能放在使用的地方定义，或者说，参数一样的方法就是相同的函数
func normal(r *http.Request, f func()) (res *Res, err error) {
	f()
	//获取头信息，token
	token := r.Header.Get("authorization")
	fmt.Println("token", token)
	//查询缓存中是否有该信息
	cookie := userStatus[token]

	if cookie.AccountId == "" {
		//查询数据库
		fmt.Println("is  empty == ", unsafe.Sizeof(cookie))
	}
	//判断市场是否有效

	//查询数据库后，还是为空，则反馈错误

	//判断是否登录状态有效,通过判断时长，确定是否有效

	//如果无效，返回异常
	return nil, errors.New(statusMsg.TOKEN_INVALID)
}
