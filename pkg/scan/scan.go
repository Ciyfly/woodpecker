/*
 * @Date: 2022-04-13 17:08:46
 * @LastEditors: recar
 * @LastEditTime: 2022-04-29 17:23:46
 */
package scan

import (
	"fmt"
	"net/http"
	"strings"
	"woodpecker/pkg/cel"
	"woodpecker/pkg/common/enum"
	"woodpecker/pkg/common/nuclei"
	"woodpecker/pkg/common/xray"
	"woodpecker/pkg/log"
	"woodpecker/pkg/parse"
	"woodpecker/pocs/scripts"
)

type ScanItem struct {
	Target     string
	XrayPoc    *xray.Poc
	NucleiPoc  *nuclei.Poc
	GoScript   scripts.ScriptFunc
	VerifyType int
	Mode       string
	TaskId     int
}

func RunXray(scanItem *ScanItem) {
	var ScanReq *http.Request
	var ScanRsp *http.Response
	var resultStatus int
	// 测试目标生成req包
	// cel-go 构建
	log.Debugf("target: %s poc: %s", scanItem.Target, scanItem.XrayPoc.Name)
	// 对于协议是tcp先不处理
	if scanItem.XrayPoc.Transport == "tcp" || scanItem.XrayPoc.Transport == "udp" {
		return
	}
	// init cel env map
	celController := &cel.CelController{}
	celController.InitCel(scanItem.XrayPoc)
	// 需要request set中需要
	pReq := parse.GetPReqByTarget(scanItem.Target)
	// 处理set
	celController.InitSet(scanItem.XrayPoc, pReq)
	// rules
	for key, rule := range scanItem.XrayPoc.Rules {
		rule.ReplaceSet(celController.ParamMap)
		// 根据原始请求 + rule 生成并发起新的请求 http
		req, rsp, prsp, err := parse.DoRequest(rule, scanItem.Target)
		ScanReq = req
		ScanRsp = rsp
		if err != nil {
			log.Errorf("XrayPoc:  %s req error: %s", scanItem.XrayPoc.Name, err.Error())
			resultStatus = enum.ResultError
			break
		}
		celController.ParamMap["response"] = prsp
		// 匹配search规则
		if rule.Output.Search != "" {
			celController.ParamMap = rule.ReplaceSearch(prsp, celController.ParamMap)
		}
		// 执行表达式
		out, err := celController.Evaluate(rule.Expression)
		if err != nil {
			log.Debugf("XrayPoc: %s rule cel evaluate error: %s", scanItem.XrayPoc.Name, err.Error())
		}
		// 将out结果写到env map里 最后再次更新env后 执行XrayPoc 表达式来判断是否成功
		// 这里更新cel的函数将rule的name的函数定义进去
		// celController.ParamMap[key] = out
		celController.UpdateRule(key, out)
	}
	// rule都跑完后要更新env 将构建的rule函数构建进去
	celController.UpdateEnv()
	out, err := celController.Evaluate(scanItem.XrayPoc.Expression)
	// 目前要求把失败的和成功都存储下来
	if out == true {
		log.Debugf("%s 漏洞存在: %s", scanItem.Target, scanItem.XrayPoc.Name)
		resultStatus = enum.ResultSuccess
	} else if resultStatus != enum.ResultError {
		resultStatus = enum.ResultFail
	}
	if err != nil {
		log.Debugf("poc: %s poc cel evaluate error: %s", scanItem.XrayPoc.Name, err.Error())
	}
	r := &Result{
		Target:       scanItem.Target,
		PocName:      scanItem.XrayPoc.Name,
		PocLink:      strings.Join(scanItem.XrayPoc.Detail.Links, ","),
		Description:  scanItem.XrayPoc.Detail.Description,
		Mode:         scanItem.Mode,
		TaskId:       scanItem.TaskId,
		ResultStatus: resultStatus,
		Req:          ScanReq,
		Rsp:          ScanRsp,
	}
	ProducerResult(r)
}

func RunGoScript(scanItem *ScanItem) {
	// 用scanItem to Script Item 然后调用函数返回结果入库
	scriptParams := &scripts.ScriptParams{
		Target: scanItem.Target,
		TaskId: scanItem.TaskId,
		Mode:   scanItem.Mode,
	}
	scriptResult := scanItem.GoScript(scriptParams)
	ProducerGoScriptResult(scriptParams, scriptResult)
}

func RunNucleic(scanItem *ScanItem) {
	var resultStatus int
	ret, _ := nuclei.ExecuteNucleiPoc(scanItem.Target, scanItem.NucleiPoc)
	resultStatus = enum.ResultFail
	if ret != nil {
		resultStatus = enum.ResultSuccess
	}
	r := &Result{
		Target:       scanItem.Target,
		PocName:      scanItem.NucleiPoc.ID,
		PocLink:      scanItem.NucleiPoc.Info.Reference.String(),
		Description:  scanItem.NucleiPoc.Info.Description,
		Mode:         scanItem.Mode,
		TaskId:       scanItem.TaskId,
		ResultStatus: resultStatus,
		// Req:          ScanReq,
		// Rsp:          ScanRsp,
	}
	ProducerResult(r)
}

func RunPoc(data interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
			EndChannel <- "Error"
		}
	}()
	scanItem := data.(*ScanItem)
	verifyType := scanItem.VerifyType
	if verifyType == enum.XrayVerify {
		RunXray(scanItem)
	} else if verifyType == enum.GoScriptVerify {
		RunGoScript(scanItem)
	} else if verifyType == enum.NuclieVerify {
		RunNucleic(scanItem)
	}

}
