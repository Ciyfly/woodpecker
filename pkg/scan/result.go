/*
 * @Date: 2022-04-14 18:02:14
 * @LastEditors: recar
 * @LastEditTime: 2022-04-14 18:15:08
 */
package scan

import "woodpecker/pkg/log"

type Result struct {
	Target      string
	PocName     string
	PocLink     string
	Description string
}

var ResultChannel chan *Result

func InitResultChannel(mode string) {
	ResultChannel = make(chan *Result)
	// 启动消费者
	if mode == "scan" {
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
		log.Infof("[+] insertdb %s -> %s", result.Target, result.PocName)
	}

}

func PrintResult() {
	for {
		result := <-ResultChannel
		log.Infof("[+] print %s -> %s", result.Target, result.PocName)
	}
}
