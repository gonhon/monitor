package config

import "monitor/monitor/internal/types"

type Monitor struct {
	Msg      []types.MessageProperties
	Mode     string
	FilePath string
	SendUrl  string
	Timeout  int
}

var m Monitor

func RegisterProperty(monitor Monitor) {
	m = monitor
}

func GetMonitor() Monitor {
	return m
}
