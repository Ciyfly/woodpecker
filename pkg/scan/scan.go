/*
 * @Date: 2022-04-13 17:08:46
 * @LastEditors: recar
 * @LastEditTime: 2022-04-14 19:31:23
 */
package scan

import (
	"strings"
	"woodpecker/pkg/cel"
	"woodpecker/pkg/log"
	"woodpecker/pkg/parse"
)

type ScanItem struct {
	Target string
	Poc    *parse.Poc
}

func RunPoc(data interface{}) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered. Error:\n", r)
	// 		EndChannel <- "Error"
	// 	}
	// }()
	// 解析执行
	scanItem := data.(*ScanItem)
	// 测试目标生成req包
	// cel-go 构建
	log.Debugf("target: %s poc: %s", scanItem.Target, scanItem.Poc.Name)
	// 对于协议是tcp先不处理
	if scanItem.Poc.Transport == "tcp" || scanItem.Poc.Transport == "udp" {
		return
	}

	// 有ip了 有 poc解析后的了 接下来是解析poc 要先生成个包? 为啥啊
	// init cel env map
	celController := &cel.CelController{}
	celController.InitCel(scanItem.Poc)
	// 需要request set中需要
	pReq := parse.GetPReqByTarget(scanItem.Target)
	// 处理set
	celController.InitSet(scanItem.Poc, pReq)
	// rules
	for key, rule := range scanItem.Poc.Rules {
		rule.ReplaceSet(celController.ParamMap)
		// 根据原始请求 + rule 生成并发起新的请求 http
		rsp, err := parse.DoRequest(rule, scanItem.Target)
		if err != nil {
			log.Errorf("poc:  %s req error: %s", scanItem.Poc.Name, err.Error())
		}
		celController.ParamMap["response"] = rsp
		// 匹配search规则
		if rule.Output.Search != "" {
			celController.ParamMap = rule.ReplaceSearch(rsp, celController.ParamMap)
		}
		// 执行表达式
		out, err := celController.Evaluate(rule.Expression)
		if err != nil {
			log.Errorf("poc: %s rule cel evaluate error: %s", scanItem.Poc.Name, err.Error())
		}
		// 将out结果写到env map里 最后再次更新env后 执行poc 表达式来判断是否成功
		celController.ParamMap[key] = out
	}
	// rule都跑完后要更新env
	out, err := celController.Evaluate(scanItem.Poc.Expression)
	if err != nil {
		log.Errorf("poc: %s poc cel evaluate error: %s", scanItem.Poc.Name, err.Error())
	}
	if out == true {
		log.Infof("poc: %s 验证成功 漏洞存在", scanItem.Poc.Name)
		r := &Result{
			Target:      scanItem.Target,
			PocName:     scanItem.Poc.Name,
			PocLink:     strings.Join(scanItem.Poc.Detail.Links, ","),
			Description: scanItem.Poc.Detail.Description,
		}
		ProducerResult(r)
	}
}
