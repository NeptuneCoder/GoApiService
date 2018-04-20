package mysessoin

import (
	"net/http"
	"status/statusMsg"
	"fmt"
	"unsafe"
	"errors"
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

//2 。TODO 需要处理不同的请求类型过来时，获取token值
func CheckSession(r *http.Request) (res *Res, err error) {
	//return normal(r, func() { //3. 这是写的时候想到的一种写法，后面发现遇到为做处理的类型，不知道怎么兼容
	//	//1 。TODO 如果这个判断放在这里面，当遇到不支持的时候，怎么处理呢？x
	//	if r.Method == "POST" {
	//		r.ParseForm()
	//	}else{
	//
	//	}
	//
	//})
	if r.Method == "POST" {
		return custom(r, func() {
			r.ParseForm()
		})
	}

	return nil, errors.New(statusMsg.REQUEST_TYPE_NO_SURPPORT)
}

//函数作为参数只能放在使用的地方定义，或者说，参数一样的方法就是相同的函数
func custom(r *http.Request, f func()) (res *Res, err error) {
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
