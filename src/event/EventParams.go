package event

type Event struct {
	Timestamp  string `json:"timestamp"`
	Event      string `json:"event"`
	TimePhone  string `json:"timePhone"`
	Uuid       string `json:"uuid"`
	Androidid  string `json:"androidid"`
	PhoneType  string `json:"phoneType"`
	Language   string `json:"language"`
	Country    string `json:"country"`
	AppVersion string `json:"appVersion"`
	OsVersion  string `json:"osVersion"`
	Segment    string `json:"segment"`
	Level      string `json:"level"`
	SdkVersion string `json:"sdkVersion"`
}
