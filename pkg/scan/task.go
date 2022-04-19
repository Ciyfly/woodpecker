/*
 * @Date: 2022-04-13 18:17:06
 * @LastEditors: recar
 * @LastEditTime: 2022-04-19 15:13:02
 */
package scan

import (
	"sync"
	"woodpecker/pkg/conf"
	"woodpecker/pkg/db"
	"woodpecker/pkg/log"
	"woodpecker/pkg/parse"

	"github.com/panjf2000/ants/v2"
)

type Task struct {
	Targets []string
	Pocs    []*parse.Poc
	Mode    string
	TaskId  int
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

func ProducerTask(mode string, targets []string, pocNames []string, pocIds []int, taskId int) {
	// 根据targets 和pocNames pocIds 推任务到task管道
	pocs := []*parse.Poc{}
	if mode == "scan" {
		// 命令行扫描
		pocs = parse.LoadYamlPoc(pocNames)
	} else {
		// 通过ids从db中获取poc并解析yaml处理
		dbPocs, err := db.GetPocByIds(pocIds)
		if err != nil {
			log.Errorf("GetPocByIds: %s", err.Error())
		}
		// db to yaml poc -> parse.Poc
		pocs = parse.DbPoc2YamlPoc(dbPocs)
	}
	task := &Task{
		Targets: targets,
		Pocs:    pocs,
		Mode:    mode,
		TaskId:  taskId,
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
	// server 模式的话要更新task 状态
	log.Info("start scan")
	if task.Mode == parse.ModeServer {
		db.UpdateTaskStatusById(task.TaskId, parse.ScanRun)
	}
	// 先验证target
	log.Info("start verify Target")
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
				Mode:   task.Mode,
				TaskId: task.TaskId,
			}
			_ = p.Invoke(scanItem)
		}
	}
	wg.Wait()
	log.Info("scan end")
	if task.Mode == parse.ModeServer {
		db.UpdateTaskStatusById(task.TaskId, parse.ScanEnd)
	} else {
		EndChannel <- "End"
	}
}
