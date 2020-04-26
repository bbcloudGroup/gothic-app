package console

import (
	"fmt"
	"github.com/bbcloudGroup/gothic/logerr"
	"gothic_app/datasource"
	"time"
)

type Printer struct {
	cache datasource.Cache
}

func NewPrinter(cache datasource.Cache) Printer {
	return Printer{cache:cache}
}

func (p Printer) Run() {

	_, err := p.cache.Incr("abc").Result()
	if err != nil {
		logerr.Error(err.Error())
		return
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " " + p.cache.Get("abc").Val())
}
