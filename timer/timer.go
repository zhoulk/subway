package timer

import (
	"time"
)

var (
	index = 1
)

func init() {
	ticker := time.NewTicker(time.Second * time.Duration(1))
	go func() {
		for range ticker.C {
			index++
			go func() {
				// beego.Informational("timer ", index)
			}()
		}
	}()
}
