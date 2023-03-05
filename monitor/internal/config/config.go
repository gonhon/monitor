package config

import (
	"monitor/monitor/internal/types"

	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Monitor struct {
		Msg      []types.MessageProperties
		Mode     string
		FilePath string
		SendUrl  string
		Timeout  int
	}
}
