package mysessoin

import (
	"testing"
	"net/http"
	"fmt"
)

func TestCheckSession(t *testing.T) {
	r := &http.Request{}
	r.Method = "GET"
	header := http.Header{}
	header.Add("authorization", "i2o3iankdakajksjfqpxpwdf")
	r.Header = header
	res, err := CheckSession(r)
	fmt.Println(res, err)
}
