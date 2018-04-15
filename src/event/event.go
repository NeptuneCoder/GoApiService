package event

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"io"
	"initialize"
	"github.com/yanghai23/GoLib/aterr"
	"github.com/yanghai23/GoLib/athttp"
	"time"
	"bytes"
	"utils"
)

type Event struct {
	Event      string
	Timestamp  string
	TimePhone  string
	Uuid       string
	Androidid  string
	PhoneType  string
	Language   string
	Country    string
	AppVersion string
	OsVersion  string
	Segment    string
	Level      string
	SdkVersion string
}

func SaveEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body) //获取post的数据
	fmt.Println(string(con))
	event := &Event{}
	json.Unmarshal(con, &event)
	SaveEvent2DB(event)
	result := make(map[string]interface{})
	result["code"] = 200
	result["msg"] = "存储成功"
	fmt.Println(string(len(result)))
	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(res))

}

func SaveEvent2DB(data *Event) {
	//插入数据
	stmt, err := initialize.Db.Prepare("INSERT INTO EventTab(event,timestamp,timePhone,uuid,androidid,phoneType,language,country,appVersion,osVersion,segment,level,sdkVersion) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	aterr.CheckErr(err)
	res, err := stmt.Exec(data.Event, data.Timestamp, data.TimePhone, data.Uuid, data.Androidid, data.PhoneType, data.Language, data.Country, data.AppVersion, data.OsVersion, data.Segment, data.Level, data.SdkVersion)
	aterr.CheckErr(err)
	_, err = res.LastInsertId()
	aterr.CheckErr(err)
}

type PayStatus struct {
	VpnId       string
	Country     string
	Version     string
	DollarPrice string
	Type        string
	ResultCode  string
	Level       string
	Code        string
	TimeStr     string
	Result      string
}

func insertPayStatusData(ps *PayStatus) {
	fmt.Println("ps.Result === === === === ", ps.Result)
	//插入数据
	stmt, err := initialize.Db.Prepare("INSERT INTO PaymentStatus(vpnId,country,version,dollarPrice,type,level,code,result,timeStr) VALUES(?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	aterr.CheckErr(err)
	fmt.Println("ps.Result === === === === ", ps.Result)
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
	ps := &PayStatus{}
	json.Unmarshal(result, &ps)
	insertPayStatusData(ps)

	athttp.HttpRequest(utils.SendNotify(initialize.BaseConfig.Url0, string(bytes.NewBuffer(result).String())))
	utils.OkStatus(w, "")
}
