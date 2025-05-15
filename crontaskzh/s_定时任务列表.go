package crontaskzh

import (
	"fmt"
	"time"

	"github.com/go-zwbc/cronzh/cronnextzh"
	"github.com/go-zwbc/timezh"
	"github.com/robfig/cron/v3"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/sortslice"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type T定时任务 struct {
	S定时表达式列表 []string           //crontab 的表达式列表(spec-list)
	E任务名称    string             //需要做的任务名称
	F执行函数    func(e任务名称 string) `json:"-"`
}

type S定时任务列表 []*T定时任务

func NewS定时任务列表(s定时任务列表 []*T定时任务) S定时任务列表 {
	return s定时任务列表
}

func (cs S定时任务列表) Set注册定时任务(cron *cron.Cron) {
	zaplog.LOG.Debug("新建定时任务")
	for idx, one := range cs {
		zaplog.LOG.Debug("注册定时任务", zap.Int("idx", idx), zap.Int("spec_list", len(one.S定时表达式列表)))
		zaplog.SUG.Debugln("定时任务列表", neatjsons.S(cs))

		for _, spec := range one.S定时表达式列表 {
			e任务名称 := one.E任务名称 //这里提取出变量是关键，否则运行和测试不正确(在 version <= go.21 时不正确)，因为里面是新线程
			f执行函数 := one.F执行函数
			zaplog.LOG.Debug("注册定时条件", zap.String("spec", spec), zap.String("event_name", e任务名称))

			rese.C1(cron.AddFunc(spec, func() {
				zaplog.LOG.Debug("开始执行", zap.String("spec", spec), zap.String("event_name", e任务名称))
				f执行函数(e任务名称)
				zaplog.LOG.Debug("执行完毕", zap.String("spec", spec), zap.String("event_name", e任务名称))
			}))
		}
	}
	zaplog.LOG.Debug("任务注册完毕")
}

func (cs S定时任务列表) Debug(p表达式解析器 *cronnextzh.P表达式解析器, nDate int) {
	type T任务和时间 struct {
		E任务名称 string
		T执行时间 time.Time
	}
	var s任务列表 []*T任务和时间
	now := time.Now()
	for _, sa := range cs {
		for _, spec := range sa.S定时表达式列表 {
			for _, v执行时间 := range p表达式解析器.Get获取未来N天内的执行时间(spec, now, nDate) {
				s任务列表 = append(s任务列表, &T任务和时间{
					E任务名称: sa.E任务名称,
					T执行时间: v执行时间,
				})
			}
		}
	}

	sortslice.SortVStable(s任务列表, func(a, b *T任务和时间) bool {
		return a.T执行时间.Before(b.T执行时间)
	})

	zaplog.SUG.Debugln("显示未来的执行时间")
	zaplog.SUG.Debugln("----------")
	for idx, tsk := range s任务列表 {
		zaplog.SUG.Debugln(
			fmt.Sprintf("%04d", idx),
			timezh.TS.T时间.Get转字符串(tsk.T执行时间),
			int(tsk.T执行时间.Weekday()),
			tsk.E任务名称,
		)
	}
	zaplog.SUG.Debugln("----------")
}
