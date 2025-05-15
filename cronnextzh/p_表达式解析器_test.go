package cronnextzh

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-zwbc/timezh"
	"github.com/gorhill/cronexpr"
)

func TestP表达式解析器_Get获取未来N天内的执行时间(t *testing.T) {
	times := P带秒数的表达式解析器.Get获取未来N天内的执行时间("30 25/10 9 * * 1-5", time.Now(), 15)
	for i, v := range times {
		fmt.Println(fmt.Sprintf("%02d", i), timezh.TS.T时间.Get转字符串(v), int(v.Weekday()))
	}
}

func TestP表达式解析器_Get获取未来N天内的执行时间_1(t *testing.T) {
	times := P带秒数的表达式解析器.Get获取未来N天内的执行时间("0 5 0,12 * * 0/6", time.Now(), 15)
	for i, v := range times {
		fmt.Println(fmt.Sprintf("%02d", i), timezh.TS.T时间.Get转字符串(v), int(v.Weekday()))
	}
}

func TestP表达式解析器_Get计算未来N天内的执行时间(t *testing.T) {
	specs := []string{
		"30 25/10 9 * * 1-5",
		"30 5/10 10,13,14 * * 1-5",
		"30 5-35/10 11 * * 1-5",
		"30 3 15 * * 1-5",
	}
	times := P带秒数的表达式解析器.Get计算未来N天内的执行时间(specs, time.Now(), 15)
	for i, v := range times {
		fmt.Println(fmt.Sprintf("%02d", i), timezh.TS.T时间.Get转字符串(v), int(v.Weekday()))
	}
}

func TestX其它工具的对照结果(t *testing.T) {
	times := cronexpr.MustParse("* * 10 * * * *").NextN(time.Now(), 5)
	for _, v := range times {
		t.Log(v)
	}
}
