package timer

import (
	"time"

	"github.com/astaxie/beego"
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
				beego.Informational("timer ", index)
			}()
		}
	}()
}
