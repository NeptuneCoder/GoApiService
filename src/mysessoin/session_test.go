package mysessoin

import (
	"testing"
	"net/http"
	"fmt"
	"time"
)
func TestGenerateCookie(t *testing.T) {

	GenerateCookie("i2o3iankdakajksjfqpxpwdf")
}

func TestCheckSession(t *testing.T) {
	r := &http.Request{}
	r.Method = "POST"
	header := http.Header{}
	header.Add("authorization", "i2o3iankdakajksjfqpxpwdf")
	r.Header = header
	res, err := CheckSession(r)
	fmt.Println(res, err)
}

func TestQueryToken(t *testing.T) {
	fmt.Println(time.Now().UnixNano())
}
