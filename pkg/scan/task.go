/*
 * @Date: 2022-04-13 18:17:06
 * @LastEditors: recar
 * @LastEditTime: 2022-04-15 17:59:08
 */
package scan

import (
	"sync"
	"woodpecker/pkg/conf"
	"woodpecker/pkg/log"
	"woodpecker/pkg/parse"

	"github.com/panjf2000/ants/v2"
)

type Task struct {
	Targets []string
	Pocs    []*parse.Poc
}

var TaskChannel chan *Task
var EndChannel chan string

func InitTaskChannel() {
	// 初始化 任务管道
	TaskChannel = make(chan *Task)
	EndChannel = make(chan string)
	// 启动消费者
	go Consumer()
}

func ProducerTask(mode string, targets []string, pocNames []string, pocIds []string) {
	// 根据targets 和pocNames pocIds 推任务到task管道
	pocs := []*parse.Poc{}
	if mode == "scan" {
		// 命令行扫描
		pocs = parse.LoadYamlPoc(pocNames)
	} else {
		// 通过ids从db中获取poc并解析yaml处理
	}
	task := &Task{
		Targets: targets,
		Pocs:    pocs,
	}
	TaskChannel <- task
}

func Consumer() {
	for {
		task := <-TaskChannel
		go RunTask(task)
	}
}

func RunTask(task *Task) {
	// 先验证target
	log.Infof("start verify Target")
	aliveTargets := VerifyTargets(task.Targets)
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(conf.GlobalConfig.RateSize, func(data interface{}) {
		RunPoc(data)
		wg.Done()
	})
	defer p.Release()
	for _, target := range aliveTargets {
		for _, poc := range task.Pocs {
			wg.Add(1)
			scanItem := &ScanItem{
				Target: target,
				Poc:    poc,
			}
			_ = p.Invoke(scanItem)
		}
	}
	wg.Wait()
	EndChannel <- "End"

}
