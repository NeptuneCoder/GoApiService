package product

import (
	"net/http"
	"fmt"
	"utils"
	"github.com/yanghai23/GoLib/aterr"
	"encoding/json"
	"mysessoin"
	"status/statusCode"
	"status/statusMsg"
	"initialize"
)

func GetOldProductInfo(w http.ResponseWriter, r *http.Request) {

}

func GetProductInfo(w http.ResponseWriter, r *http.Request) {
	_, err := mysessoin.CheckSession(r)
	if err != nil {
		utils.RStatus(w, statusCode.LOGIN_INVALID, statusMsg.LOGIN_INVALID, "")
		http.RedirectHandler("", 200)
		return
	}
	service := QueryService()
	buf, err := json.Marshal(service)
	aterr.CheckErr(err)
	utils.RStatus(w, statusCode.SUCCESS, statusMsg.SERVICE_TYPE, string(buf))
}

func QueryService() (service []ServiceTypeParam) {
	//获取登陆token，验证用户是否合法
	//获取用户名，判断是否已购买过商品
	//获取所在国家
	rows, err := initialize.Db.Query("SELECT serviceType,ad,video,speed,image,serviceExplain FROM ServiceTypeTab")
	defer rows.Close()
	if err != nil {
		panic(statusMsg.CREATE_SQL_OBJ_ERROR)
	}
	for rows.Next() {
		data := ServiceTypeParam{}
		err = rows.Scan(&data.ServiceType, &data.Ad, &data.Video, &data.Speed, &data.Image, &data.ServiceExplain)
		aterr.CheckErr(err)
		service = append(service, data)
	}
	return
}
func GenerateOrder(w http.ResponseWriter, r *http.Request) {
	_, err := mysessoin.CheckSession(r)
	if err != nil {
		utils.RStatus(w, statusCode.LOGIN_INVALID, statusMsg.LOGIN_INVALID, "")
		return
	}
	utils.RStatus(w, statusCode.SUCCESS, statusMsg.GERERNATE_ORDER_SUCCESS, "{\"key\":\"\"}")
}

func PaymentInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Header.Get("language"))
	fmt.Println("value:", r.Form.Get("key"))
	utils.RStatus(w, statusCode.SUCCESS, "支付成功", "{\"key\":\"\"}")
}
