package mysessoin

import (
	"net/http"
	"status/statusMsg"
	"fmt"
	"unsafe"
	"errors"
	"initialize"
	"utils"
	"time"
)

var userStatus = make(map[string]*Cookie)



func (this Cookie) isExist() bool {
	return this.AccountId != ""
}
func (this Cookie) updateValidTime() {
	this.ValidTime = 24 * 10 * 60 * 60 * time.Second
}

func (this Cookie) isValid() bool {
	return this.CurTime+int64(this.ValidTime) > time.Now().UnixNano()
}

type Res struct {
	ErrMsg  string
	ErrCode int
}

func GenerateCookie(accountId string) (error) {
	//gernate token

	defer func() error {
		if e, ok := recover().(error); ok {
			return e
		}
		return nil
	}()
	cookie := &Cookie{AccountId: accountId,                     //
		Token: string(utils.Krand(28, utils.KC_RAND_KIND_ALL)), //
		CurTime: time.Now().UnixNano(),                         //
		ValidTime: 24 * 10 * 60 * 60 * time.Second}

	userStatus[accountId] = cookie
	fmt.Println(cookie.Token)

	return nil
}

//2 。TODO 需要处理不同的请求类型过来时，获取token值
func CheckSession(r *http.Request) (res *Res, err error) {
	//return normal(r, func() { //3. 这是写的时候想到的一种写法，后面发现遇到为做处理的类型，不知道怎么兼容
	//	//1 。TODO 如果这个判断放在这里面，当遇到不支持的时候，怎么处理呢？x
	//	if r.Method == "POST" {
	//		r.ParseForm()
	//	}else{
	//	}
	//})
	if r.Method == "POST" {
		return custom(r, func() {
			r.ParseForm()
		})
	} else if r.Method == "GET" {

	}

	return nil, errors.New(statusMsg.REQUEST_TYPE_NO_SURPPORT)
}

//函数作为参数只能放在使用的地方定义，或者说，参数一样的方法就是相同的函数
//第一个参数是请求信息
//第二个参数是决定怎么解析请求中的数据
//第三个参数是解决怎么实现数据库查询功能
func custom(r *http.Request, f func()) (res *Res, err error) {
	f()
	//获取头信息，token
	token := r.Header.Get("authorization")
	fmt.Println("token", token)
	//查询缓存中是否有该信息
	cookie := userStatus[token]

	if !cookie.isExist() {
		//查询数据库
		fmt.Println("is  empty == ", unsafe.Sizeof(cookie))
		//TODO 数据库的先不测试
		//cookie, err = QueryTokenByDb(token)
	}

	if cookie.isValid() {
		cookie.updateValidTime()
		return nil, nil
	}

	//判断市场是否有效

	//查询数据库后，还是为空，则反馈错误

	//判断是否登录状态有效,通过判断时长，确定是否有效

	//如果无效，返回异常
	//4 TODO 指针类型或者是接口类型才可以将nil作为默认值，将其作为返回值，
	//5 换句话说，结构体直接就是实体，不存在nil的情况，结构体内的字段信息是默认值
	return nil, errors.New(statusMsg.TOKEN_INVALID)
}

func QueryTokenByDb(token string) (res *Cookie, err error) {
	ck := &Cookie{Token: token}
	rows, err := initialize.Db.Query("SELECT accountId,time,validTime FROM CookieTab WHERE token = ?", token)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(ck.AccountId, ck.CurTime, ck.ValidTime)
	}
	return ck, nil
}
