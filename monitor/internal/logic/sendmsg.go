package logic

import (
	"monitor/monitor/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

func SendMsg(msgs []types.MessageProperties) error {
	logx.Info(msgs)

	return nil
}
