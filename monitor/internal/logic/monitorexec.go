package logic

import (
	"fmt"
	"monitor/monitor/internal/svc"
	"monitor/monitor/internal/types"

	"github.com/robfig/cron"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	LocalStrategy    = "local"
	PropertyStrategy = "property"
)

var (
	localMsgStrategyBuilder    localMsgStrategy
	propertyMsgStrategyBuilder propertyMsgStrategy
	m                          = make(map[string]LoadMsgStrategy)
)

// 加载数据
type LoadMsgStrategy interface {
	loadData() ([]types.MessageProperties, error)
	Scheme() string
}

func Register(b LoadMsgStrategy) {
	m[b.Scheme()] = b
}

func init() {
	Register(&localMsgStrategyBuilder)
	Register(&propertyMsgStrategyBuilder)

}

func RegisterSchedule(serverCtx *svc.ServiceContext) {

	mode := serverCtx.Config.Monitor.Mode

	strategy, ok := m[mode]
	if !ok {
		logx.Error("")
	}
	msgs, err := strategy.loadData()
	if err != nil {
		logx.Error(err)
	}
	crons := cron.New()

	if err := crons.AddFunc("*/30 * * * * *", func() {
		SendMsg(msgs)
	}); err != nil {
		fmt.Println(err)
	}
	crons.Start()
}

func example() {
	logx.Info("Hello World!")
}
