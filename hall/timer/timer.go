package timer

import (
	"subway/hall/models"
	"time"
)

var (
	index = 1
)

func init() {
	ticker := time.NewTicker(time.Second * time.Duration(5))
	go func() {
		for range ticker.C {
			index++
			go func() {
				// beego.Informational("timer ", index)

				models.Persistent()
			}()
		}
	}()
}
