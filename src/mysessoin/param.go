package mysessoin

import "time"

type CookieParam struct {
	AccountId string
	Token     string
	CurTime   int64
	ValidTime time.Duration
}
