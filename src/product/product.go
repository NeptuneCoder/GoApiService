package product

import (
	"net/http"
	"fmt"
	"utils"
	"github.com/yanghai23/GoLib/aterr"
	"initialize"
	"encoding/json"
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
}

func GetProductInfo(w http.ResponseWriter, r *http.Request) {
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
	utils.OkStatus(w, 200, "服务类型", string(buf))
}

func GenerateOrder(w http.ResponseWriter, r *http.Request) {

	utils.OkStatus(w, 200, "支付成功", "{\"key\":\"\"}")
}

func PaymentInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Header.Get("language"))
	fmt.Println("value:", r.Form.Get("key"))

	utils.OkStatus(w, 200, "支付成功", "{\"key\":\"\"}")
}
