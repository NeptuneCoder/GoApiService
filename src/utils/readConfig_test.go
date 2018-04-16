package utils

import (
	"testing"
	"fmt"
	"github.com/yanghai23/GoLib/atfile"
)

func TestReadBaseConfig(t *testing.T) {
	fmt.Println(atfile.GetCurrentDirectory())
	data, _ := ReadBaseConfig(atfile.GetCurrentDirectory() + "/Config.json")
	fmt.Println("url = ",data.BgUrl)
}
