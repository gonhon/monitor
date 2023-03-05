package logic

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/zeromicro/go-zero/core/logx"

	"monitor/monitor/internal/types"
)

type localMsgStrategy struct{}

func (l *localMsgStrategy) loadData() ([]types.MessageProperties, error) {
	jsonFile, err := os.Open("D:\\project\\go_project\\monitor\\source.json")
	// 最好要处理以下错误
	if err != nil {
		logx.Error(err)
	}

	// 要记得关闭
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var msgs []types.MessageProperties
	json.Unmarshal([]byte(byteValue), &msgs)
	return msgs, nil
}

func (l *localMsgStrategy) Scheme() string {
	return LocalStrategy
}
