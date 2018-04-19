package product

import (
	"net/http"
	"fmt"
	"utils"
	"github.com/yanghai23/GoLib/aterr"
	"initialize"
	"encoding/json"

	"mysessoin"
	"status/statusCode"
	"status/statusMsg"
)

func GetOldProductInfo(w http.ResponseWriter, r *http.Request) {

}

type ServiceType struct {
	ServiceType    string `josn:"serviceType"`
	Ad             string `josn:"ad"`
	Video          string `josn:"video"`
	Speed          string `josn:"speed"`
	Image          string `josn:"image"`
	ServiceExplain string `josn:"serviceExplain"`
	Products       []Product
}

type Product struct {
	serviceType string
	productId   string
	price       string
	priceUnit   string
	validTime   string
	timeUnit    string
}

func GetProductInfo(w http.ResponseWriter, r *http.Request) {
	err := mysessoin.CheckSession(r)
	if err != nil {
		utils.OkStatus(w, statusCode.LOGIN_INVALID, statusMsg.LOGIN_INVALID, "")
		return
	}

	//获取登陆token，验证用户是否合法
	//获取用户名，判断是否已购买过商品
	//获取所在国家
	rows, err := initialize.Db.Query("SELECT serviceType,ad,video,speed,image,serviceExplain FROM ServiceTypeTab WHERE accountId = ?", accountId)
	defer rows.Close()
	var service []ServiceType
	if err != nil {
		panic("创建sql句柄错误")
	}
	for rows.Next() {
		data := ServiceType{}
		err = rows.Scan(&data.ServiceType, &data.Ad, &data.Video, &data.Speed, &data.Image, &data.ServiceExplain)
		aterr.CheckErr(err)
		service = append(service, data)
	}
	buf, err := json.Marshal(service)
	aterr.CheckErr(err)
	utils.OkStatus(w, statusCode.SUCCESS, "服务类型", string(buf))
}

func GenerateOrder(w http.ResponseWriter, r *http.Request) {

	utils.OkStatus(w, statusCode.SUCCESS, "支付成功", "{\"key\":\"\"}")
}

func PaymentInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Header.Get("language"))
	fmt.Println("value:", r.Form.Get("key"))

	utils.OkStatus(w, statusCode.SUCCESS, "支付成功", "{\"key\":\"\"}")
}
