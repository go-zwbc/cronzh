package crontaskzh

import (
	"testing"
	"time"

	"github.com/go-zwbc/cronzh/cronnextzh"
	cronv3 "github.com/robfig/cron/v3"
	"github.com/yyle88/must"
)

func TestNewS定时任务列表_5D(t *testing.T) {
	run := func(e事件名称 string) {
		t.Log(e事件名称)
	}

	s定时任务列表 := NewS定时任务列表([]*T定时任务{
		&T定时任务{
			E任务名称:    "XXX",
			S定时表达式列表: []string{"0 30 8,20 * * 1-5"},
			F执行函数:    run,
		},
		&T定时任务{
			E任务名称:    "YYY",
			S定时表达式列表: []string{"0 0,30 7,21,22 * * 1-5"},
			F执行函数:    run,
		},
		&T定时任务{
			E任务名称:    "ZZZ",
			S定时表达式列表: []string{"0 0,30 10,11,12,13 * * 1-5"},
			F执行函数:    run,
		},
	})
	s定时任务列表.Debug(cronnextzh.P带秒数的表达式解析器, 10)
}

func TestNewS定时任务列表_6D(t *testing.T) {
	run := func(e事件名称 string) {
		t.Log(e事件名称)
	}

	s定时任务列表 := NewS定时任务列表([]*T定时任务{
		&T定时任务{
			E任务名称:    "AAA",
			S定时表达式列表: []string{"30 8,20 * * 1-5"},
			F执行函数:    run,
		},
		&T定时任务{
			E任务名称:    "BBB",
			S定时表达式列表: []string{"0,30 7,21,22 * * 1-5"},
			F执行函数:    run,
		},
		&T定时任务{
			E任务名称:    "CCC",
			S定时表达式列表: []string{"0,30 10,11,12,13 * * 1-5"},
			F执行函数:    run,
		},
	})
	s定时任务列表.Debug(cronnextzh.P只到分的表达式解析器, 10)
}

func TestS定时任务列表_Debug(t *testing.T) {
	v定时任务 := &T定时任务{
		S定时表达式列表: []string{
			"0 5 0,8,9,12,21 * * 1-5", //工作日
			"0 5 0,12 * * 0/6",        //休息日
		},
		E任务名称: "WWW",
		F执行函数: func(e事件名称 string) {
			t.Log(e事件名称)
		},
	}
	s定时任务列表 := NewS定时任务列表([]*T定时任务{v定时任务})
	s定时任务列表.Debug(cronnextzh.P带秒数的表达式解析器, 10)
}

func TestS定时任务列表_Set注册定时任务(t *testing.T) {
	run := func(e事件名称 string) {
		t.Log(e事件名称)
		switch e事件名称 {
		case "哈哈":
			t.Log(e事件名称 + "->哈哈哈")
		case "嘿嘿":
			t.Log(e事件名称 + "->嘿嘿嘿")
		}
	}

	s定时任务列表 := NewS定时任务列表([]*T定时任务{
		&T定时任务{
			E任务名称:    "哈哈",
			S定时表达式列表: []string{"*/1 * * * * *", "*/3 * * * * *"},
			F执行函数:    run,
		},
		&T定时任务{
			E任务名称:    "嘿嘿",
			S定时表达式列表: []string{"*/2 * * * * ?"},
			F执行函数:    run,
		},
	})
	must.Have(s定时任务列表)

	cron := cronv3.New(cronv3.WithSeconds(), cronv3.WithLocation(time.UTC))
	s定时任务列表.Set注册定时任务(cron)
	cron.Start()

	time.Sleep(time.Second * 5)
	ctx := cron.Stop()
	<-ctx.Done()
}
