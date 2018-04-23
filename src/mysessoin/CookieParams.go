package mysessoin

import "time"

type Cookie struct {
	AccountId string
	Token     string
	CurTime   int64
	ValidTime time.Duration
}
