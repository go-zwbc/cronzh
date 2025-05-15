package cronnextzh

import (
	"time"

	"github.com/go-zwbc/timezh"
	"github.com/robfig/cron/v3"
	"github.com/yyle88/rese"
	"github.com/yyle88/sortslice"
)

var P带秒数的表达式解析器 = New(cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor))
var P只到分的表达式解析器 = New(cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor))

type P表达式解析器 cron.Parser

func New(parser cron.Parser) *P表达式解析器 {
	return (*P表达式解析器)(&parser)
}

func (P *P表达式解析器) Get获取未来N天内的执行时间(spec string, since time.Time, nDate int) []time.Time {
	p解析器 := (*cron.Parser)(P)
	s时刻表 := rese.V1(p解析器.Parse(spec))
	v执行时间 := since
	e结束时间 := timezh.TS.D日期.Get转字符串(since.AddDate(0, 0, nDate))
	var result []time.Time
	for {
		// 使用这个也是可以的 schedule := cronexpr.MustParse("* * 10 * * * *").NextN(time.Now(), 5) // 这是其它思路
		v执行时间 = s时刻表.Next(v执行时间)
		if timezh.TS.D日期.Get转字符串(v执行时间) > e结束时间 {
			return result
		}
		result = append(result, v执行时间)
	}
}

func (P *P表达式解析器) Get计算未来N天内的执行时间(specs []string, since time.Time, nDate int) []time.Time {
	var result []time.Time
	for _, spec := range specs {
		res := P.Get获取未来N天内的执行时间(spec, since, nDate)
		result = append(result, res...)

	}
	sortslice.SortVStable(result, func(a, b time.Time) bool {
		return a.Before(b)
	})
	return result
}
