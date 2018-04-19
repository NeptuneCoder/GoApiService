package utils

import (
	"testing"
	"github.com/yanghai23/GoLib/athttp"
	"github.com/yanghai23/GoLib/atfile"
	"fmt"
)

func TestSendNotify(t *testing.T) {
	fmt.Println(atfile.GetCurrentDirectory())
	data, _ := ReadBaseConfig(atfile.GetCurrentDirectory() + "/Config.json")
	fmt.Println("url = ", data.BgUrl)
	athttp.HttpRequest(SendNotify(data.LogRebootUrl, "unit test"))
}
