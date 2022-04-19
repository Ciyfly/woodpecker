/*
 * @Date: 2022-04-13 17:08:46
 * @LastEditors: recar
 * @LastEditTime: 2022-04-19 10:46:46
 */
package scan

import (
	"fmt"
	"net/http"
	"strings"
	"woodpecker/pkg/cel"
	"woodpecker/pkg/log"
	"woodpecker/pkg/parse"
)

type ScanItem struct {
	Target string
	Poc    *parse.Poc
	Mode   string
	TaskId int
}

func RunPoc(data interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
			EndChannel <- "Error"
		}
	}()
	// 解析执行
	scanItem := data.(*ScanItem)
	var ScanReq *http.Request
	var ScanRsp *http.Response
	// 测试目标生成req包
	// cel-go 构建
	log.Debugf("target: %s poc: %s", scanItem.Target, scanItem.Poc.Name)
	// 对于协议是tcp先不处理
	if scanItem.Poc.Transport == "tcp" || scanItem.Poc.Transport == "udp" {
		return
	}
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
		req, rsp, prsp, err := parse.DoRequest(rule, scanItem.Target)
		ScanReq = req
		ScanRsp = rsp
		if err != nil {
			log.Errorf("poc:  %s req error: %s", scanItem.Poc.Name, err.Error())
		}
		celController.ParamMap["response"] = prsp
		// 匹配search规则
		if rule.Output.Search != "" {
			celController.ParamMap = rule.ReplaceSearch(prsp, celController.ParamMap)
		}
		// 执行表达式
		out, err := celController.Evaluate(rule.Expression)
		if err != nil {
			log.Debugf("poc: %s rule cel evaluate error: %s", scanItem.Poc.Name, err.Error())
		}
		// 将out结果写到env map里 最后再次更新env后 执行poc 表达式来判断是否成功
		// 这里更新cel的函数将rule的name的函数定义进去
		// celController.ParamMap[key] = out
		celController.UpdateRule(key, out)
	}
	// rule都跑完后要更新env 将构建的rule函数构建进去
	celController.UpdateEnv()
	out, err := celController.Evaluate(scanItem.Poc.Expression)
	// 目前要求把失败的和成功都存储下来
	var resultStatus int
	if out == true {
		log.Debugf("%s 漏洞存在: %s", scanItem.Target, scanItem.Poc.Name)
		resultStatus = parse.ResultSuccess
	} else {
		resultStatus = parse.ResultFail
	}
	if err != nil {
		log.Debugf("poc: %s poc cel evaluate error: %s", scanItem.Poc.Name, err.Error())
	}
	r := &Result{
		Target:       scanItem.Target,
		PocName:      scanItem.Poc.Name,
		PocLink:      strings.Join(scanItem.Poc.Detail.Links, ","),
		Description:  scanItem.Poc.Detail.Description,
		Mode:         scanItem.Mode,
		TaskId:       scanItem.TaskId,
		ResultStatus: resultStatus,
		Req:          ScanReq,
		Rsp:          ScanRsp,
	}
	ProducerResult(r)
}
