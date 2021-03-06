/*
 * @Date: 2022-04-14 18:02:14
 * @LastEditors: recar
 * @LastEditTime: 2022-04-27 16:21:50
 */
package scan

import (
	"net/http"

	"woodpecker/pkg/common/enum"
	"woodpecker/pkg/db"
	"woodpecker/pkg/log"
	"woodpecker/pkg/util"
	"woodpecker/pocs/scripts"
)

type Result struct {
	Target       string
	PocName      string
	PocLink      string
	Description  string
	Mode         string
	TaskId       int
	ResultStatus int
	Req          *http.Request
	Rsp          *http.Response
}

var ResultChannel chan *Result

func InitResultChannel(mode string) {
	ResultChannel = make(chan *Result)
	if mode == enum.ModeScan {
		go PrintResult()
	} else {
		go InsertDB()
	}
}

func ProducerResult(result *Result) {
	ResultChannel <- result
}

func InsertDB() {
	for {
		result := <-ResultChannel
		dbPoc, err := db.GetPocByName(result.PocName)
		if err != nil {
			log.Errorf("GetPocByName error: %s", err.Error())
		}
		report := &db.Report{
			TaskId: result.TaskId,
			PocId:  int(dbPoc.Id),
			Target: result.Target,
			Status: result.ResultStatus,
			Req:    util.Req2Text(result.Req),
			Rsp:    util.Rsp2Text(result.Rsp),
		}
		dbResult := db.AddReport(report)
		if dbResult.Error != nil {
			log.Errorf("add Report error: %s", dbResult.Error.Error())
		}
		// 更新已经测试的数量
		// testNumber, err := db.GetTaskTestNumberById(result.TaskId)
		// if err != nil {
		// log.Errorf("testNumber error%s  %s -> %s", result.Target, result.TaskId, result.PocName)
		// }
		// db.UpdateTaskTestNumberById(result.TaskId, testNumber+1)
		db.UpdateTaskTestNumberById(result.TaskId)

	}

}

func PrintResult() {
	for {
		result := <-ResultChannel
		if result.ResultStatus == enum.ResultSuccess {
			log.Infof("[+] print %s -> %s", result.Target, result.PocName)
		}
	}
}

// go script
func ProducerGoScriptResult(sp *scripts.ScriptParams, scriptResult bool) {
	var resultStatus int
	if scriptResult == true {
		resultStatus = enum.ResultSuccess
	} else {
		resultStatus = enum.ResultFail
	}
	result := &Result{
		Target:       sp.Target,
		PocName:      sp.PocName,
		PocLink:      sp.PocLink,
		Description:  sp.Description,
		Mode:         sp.Mode,
		TaskId:       sp.TaskId,
		ResultStatus: resultStatus,
		Req:          sp.Req,
		Rsp:          sp.Rsp,
	}
	ProducerResult(result)
}
