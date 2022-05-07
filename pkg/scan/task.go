/*
 * @Date: 2022-04-13 18:17:06
 * @LastEditors: recar
 * @LastEditTime: 2022-04-29 17:20:52
 */
package scan

import (
	"sync"
	"woodpecker/pkg/common/enum"
	"woodpecker/pkg/common/nuclei"
	"woodpecker/pkg/common/xray"
	"woodpecker/pkg/conf"
	"woodpecker/pkg/db"
	"woodpecker/pkg/log"
	"woodpecker/pocs/scripts"

	"github.com/panjf2000/ants/v2"
)

type Task struct {
	Targets    []string
	XrayPocs   []*xray.Poc
	NucleiPocs []*nuclei.Poc
	GoScripts  []scripts.ScriptFunc
	Mode       string
	TaskId     int
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
	xrayPocs := []*xray.Poc{}
	nucleiPocs := []*nuclei.Poc{}
	goScripts := []scripts.ScriptFunc{}
	if mode == "scan" {
		// 命令行扫描
		// 加载poc
		xrayPocs = xray.LoadXrayYamlPoc(pocNames)
		nucleiPocs = nuclei.LoadNucleiYamlPoc(pocNames)
		goScripts = scripts.LoadScriptPoc(pocNames)
		log.Infof("xrayPoc Count: %d", len(xrayPocs))
		log.Infof("nucleiPoc Count: %d", len(nucleiPocs))
		log.Infof("goScripts Count: %d", len(goScripts))

	} else {
		// 通过ids从db中获取poc并解析yaml处理
		dbPocs, err := db.GetPocByIds(pocIds)
		if err != nil {
			log.Errorf("GetPocByIds: %s", err.Error())
		}
		xrayPocs = xray.DbPoc2XrayPoc(dbPocs)
		nucleiPocs = nuclei.DbPoc2NucleiPoc(dbPocs)
		goScripts = scripts.Db2ScriptPoc(dbPocs)

	}
	task := &Task{
		Targets:    targets,
		XrayPocs:   xrayPocs,
		NucleiPocs: nucleiPocs,
		GoScripts:  goScripts,
		Mode:       mode,
		TaskId:     taskId,
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
	if task.Mode == enum.ModeServer {
		db.UpdateTaskStatusById(task.TaskId, enum.ScanRun)
	}
	// 先验证target
	log.Info("start verify Target")
	aliveTargets := VerifyTargets(task.Targets)
	if len(aliveTargets) == 0 {
		log.Errorf("all target not connect")
	}
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(conf.GlobalConfig.RateSize, func(data interface{}) {
		RunPoc(data)
		wg.Done()
	})
	defer p.Release()
	// 计算总计需要多少 目标*poc数量
	totalNumber := len(aliveTargets) * len(task.XrayPocs)
	totalNumber += len(aliveTargets) * len(task.NucleiPocs)
	totalNumber += len(aliveTargets) * len(task.GoScripts)
	if task.Mode == enum.ModeServer {
		db.UpdateTaskTotalNumberById(task.TaskId, totalNumber)
	}
	for _, target := range aliveTargets {
		// xray
		for _, xrayPoc := range task.XrayPocs {
			wg.Add(1)
			scanItem := &ScanItem{
				Target:     target,
				XrayPoc:    xrayPoc,
				VerifyType: enum.XrayVerify,
				Mode:       task.Mode,
				TaskId:     task.TaskId,
			}
			_ = p.Invoke(scanItem)
		}
		// go script
		for _, goScript := range task.GoScripts {
			wg.Add(1)
			scanItem := &ScanItem{
				Target:     target,
				GoScript:   goScript,
				VerifyType: enum.GoScriptVerify,
				Mode:       task.Mode,
				TaskId:     task.TaskId,
			}
			_ = p.Invoke(scanItem)
		}
		// nuclei
		for _, nucleiPoc := range task.NucleiPocs {
			wg.Add(1)
			scanItem := &ScanItem{
				Target:     target,
				NucleiPoc:  nucleiPoc,
				VerifyType: enum.NuclieVerify,
				Mode:       task.Mode,
				TaskId:     task.TaskId,
			}
			_ = p.Invoke(scanItem)
		}

	}
	wg.Wait()
	if task.Mode == enum.ModeServer {
		log.Info("scan end")
		db.UpdateTaskStatusById(task.TaskId, enum.ScanEnd)
	} else {
		EndChannel <- "End"
	}
}
