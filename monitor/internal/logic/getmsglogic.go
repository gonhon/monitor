package logic

import (
	"context"

	"monitor/monitor/internal/svc"
	"monitor/monitor/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMsgLogic {
	return &GetMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMsgLogic) GetMsg() (resp []types.MessageProperties, err error) {
	// todo: add your logic here and delete this line
	resp = l.svcCtx.Config.Monitor.Msg
	return
}
