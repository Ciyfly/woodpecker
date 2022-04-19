/*
 * @Date: 2022-04-14 18:02:14
 * @LastEditors: recar
 * @LastEditTime: 2022-04-19 15:39:45
 */
package scan

import (
	"net/http"

	"woodpecker/pkg/db"
	"woodpecker/pkg/log"
	"woodpecker/pkg/parse"
	"woodpecker/pkg/util"
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
	if mode == parse.ModeScan {
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
	}

}

func PrintResult() {
	for {
		result := <-ResultChannel
		if result.ResultStatus == parse.ResultSuccess {
			log.Infof("[+] print %s -> %s", result.Target, result.PocName)
		}
	}
}
