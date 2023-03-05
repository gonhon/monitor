package logic

import (
	"monitor/monitor/internal/types"
)

type propertyMsgStrategy struct{}

func (p *propertyMsgStrategy) loadData() ([]types.MessageProperties, error) {
	return nil, nil
}

func (p *propertyMsgStrategy) Scheme() string {
	return PropertyStrategy
}
