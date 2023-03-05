package logic

import (
	"monitor/monitor/internal/config"
	"monitor/monitor/internal/types"
)

type propertyMsgStrategy struct {
}

func (p *propertyMsgStrategy) LoadData() ([]types.MessageProperties, error) {
	return config.GetMonitor().Msg, nil
}

func (p *propertyMsgStrategy) Scheme() string {
	return PropertyStrategy
}
