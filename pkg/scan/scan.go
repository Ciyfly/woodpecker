/*
 * @Date: 2022-04-13 17:08:46
 * @LastEditors: recar
 * @LastEditTime: 2022-04-14 14:55:23
 */
package scan

import (
	"woodpecker/pkg/log"
	"woodpecker/pkg/parse"
)

type ScanItem struct {
	Target string
	Poc    *parse.Poc
}

func RunPoc(data interface{}) {
	// 解析执行
	scanItem := data.(*ScanItem)
	// 测试目标生成req包
	// cel-go 构建
	log.Infof("target: %s poc: %s", scanItem.Target, scanItem.Poc.Name)
	// 有ip了 有 poc解析后的了 接下来是解析poc 要先生成个包? 为啥啊

}
