package logic

import (
	"monitor/monitor/internal/types"
	"net/http"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var cache sync.Map

func SendMsg(msgs []types.MessageProperties) error {
	// logx.Info(msgs)
	for _, m := range msgs {
		go doSendMsg(m)
	}

	return nil
}

func doSendMsg(msg types.MessageProperties) {
	resp, err := http.Get(msg.Url)
	if err != nil {
		logx.Error(err)

		now := time.Now()
		data, _ := cache.LoadOrStore(msg.Url, &cacheCycle{
			Url:       msg.Url,
			FirstTime: now,
			NextTime:  now.Add(time.Minute * 1),
			Cycle:     1,
			FirstExec: true,
		})
		logx.Info(data)
		cc := data.(*cacheCycle)
		if cc.FirstExec {
			cc.FirstExec = false
			pushMsg(msg)
		} else {
			now = time.Now()
			format := "2006-01-02 15:04:05.000 Mon Jan"
			logx.Info("当前时间:" + now.Format(format) + ",下次执行时间:" + cc.NextTime.Format(format))
			if now.After(cc.NextTime) {
				cc.Cycle <<= 1
				cc.NextTime = now.Add(time.Minute * time.Duration(cc.Cycle))
				// cache.Store(msg.Url, cc)
				go pushMsg(msg)
			} else {
				logx.Info("还未到执行时间")
			}
		}
		return
	} else {
		logx.Info(msg.Name + "-响应成功")
		if _, ok := cache.LoadAndDelete(msg.Url); ok {
			logx.Info("服务恢复...")
		}
	}
	defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// logx.Info(string(body))
}

type cacheCycle struct {
	Url string
	//首次执行时间
	FirstTime time.Time
	//下次执行时间
	NextTime time.Time
	//周期
	Cycle     int
	FirstExec bool
}

func pushMsg(msg types.MessageProperties) {
	logx.Info(msg.Name + "-准备推送消息...")
}
