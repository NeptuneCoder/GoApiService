package event

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"initialize"
	"github.com/yanghai23/GoLib/aterr"
	"github.com/yanghai23/GoLib/athttp"
	"time"
	"bytes"
	"utils"
	"status/statusCode"
	"status/statusMsg"
)

func SaveEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body) //获取post的数据
	event := &EventParam{}
	json.Unmarshal(con, &event)
	SaveEvent2DB(event)
	utils.RStatus(w, statusCode.SUCCESS, statusMsg.SAVE_SUCCES, "")

}

func SaveEvent2DB(data *EventParam) {
	//插入数据
	stmt, err := initialize.Db.Prepare("INSERT INTO EventTab(event,timestamp,timePhone,uuid,androidid,phoneType,language,country,appVersion,osVersion,segment,level,sdkVersion) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	aterr.CheckErr(err)
	res, err := stmt.Exec(data.Event, data.Timestamp, data.TimePhone, data.Uuid, data.AndroidId, data.PhoneType, data.Language, data.Country, data.AppVersion, data.OsVersion, data.Segment, data.Level, data.SdkVersion)
	aterr.CheckErr(err)
	_, err = res.LastInsertId()
	aterr.CheckErr(err)
}

func savePayStatusData(ps *PayStatusParam) {
	//插入数据
	stmt, err := initialize.Db.Prepare("INSERT INTO PaymentStatus(vpnId,country,version,dollarPrice,type,level,code,result,timeStr) VALUES(?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	aterr.CheckErr(err)
	curTime := time.Now().Format("2006-01-02")

	res, err := stmt.Exec(ps.VpnId, ps.Country, ps.Version, ps.DollarPrice, ps.Type, ps.Level, ps.Code, ps.Result, curTime)
	aterr.CheckErr(err)
	_, err = res.LastInsertId()
	aterr.CheckErr(err)
}

func PaymentStatus(w http.ResponseWriter, r *http.Request) {
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	ps := &PayStatusParam{}
	json.Unmarshal(result, &ps)
	savePayStatusData(ps)

	athttp.HttpRequest(utils.SendNotify(initialize.BaseConfig.PayRebootUrl, string(bytes.NewBuffer(result).String())))
	utils.RStatus(w, statusCode.SUCCESS, statusMsg.SUBMIT_SUCCES, "")
}
