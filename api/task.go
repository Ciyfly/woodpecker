/*
 * @Date: 2022-04-20 09:59:35
 * @LastEditors: recar
 * @LastEditTime: 2022-04-27 15:23:24
 */
package api

import (
	"fmt"
	"strconv"
	"strings"
	"woodpecker/pkg/common/enum"
	"woodpecker/pkg/db"
	"woodpecker/pkg/log"
	"woodpecker/pkg/parse"
	"woodpecker/pkg/scan"

	"github.com/gin-gonic/gin"
)

type ReqTask struct {
	TaskId   int    `json:"task_id"`
	PocIds   []int  `json:"poc_ids"`
	Port     string `json:"port"`
	Targets  string `json:"targets"`
	TaskDesc string `json:"task_desc"`
	TaskName string `json:"task_name"`
}

func AddTask(c *gin.Context) {
	reqTask := ReqTask{}
	err := c.ShouldBindJSON(&reqTask)
	if err != nil {
		c.JSON(ErrResp(err.Error()))
		return
	}
	// 先判断名字是否重复
	_, err = db.GetTaskByName(reqTask.TaskName)
	if err == nil {
		c.JSON(ErrResp("任务名称重复"))
		return
	}
	task := &db.Task{
		TaskName: reqTask.TaskName,
		TaskDesc: reqTask.TaskDesc,
		Target:   reqTask.Targets,
		Port:     reqTask.Port,
		PocIds:   strings.Trim(strings.Join(strings.Fields(fmt.Sprint(reqTask.PocIds)), ","), "[]"),
	}
	result := db.AddTask(task)
	if result.Error != nil {
		c.JSON(ErrResp(result.Error.Error()))
		return
	} else {
		targets, err := parse.ParseTargets(reqTask.Targets, reqTask.Port)
		if err != nil {
			c.JSON(ErrResp(err.Error()))
			return
		}
		scan.ProducerTask(enum.ModeServer, targets, nil, reqTask.PocIds, task.Id)
		c.JSON(SuccessResp(task))
		return
	}
}

func DelTask(c *gin.Context) {
	reqTask := &ReqTask{}
	err := c.ShouldBindJSON(&reqTask)
	if err != nil {
		c.JSON(ErrResp(err.Error()))
		return
	}
	db.DelTaskById(reqTask.TaskId)
}

func StartTask(c *gin.Context) {
	c.JSON(SuccessResp("start"))
	return
}

func StopTask(c *gin.Context) {
	c.JSON(SuccessResp("stop"))
	return
}

func GetTaskInfo(c *gin.Context) {
	log.Info("GetTaskInfo")
	taskId := c.Param("id")
	log.Infof("taskId: %s", taskId)
	taskIdNum, err := strconv.Atoi(taskId)
	if err != nil {
		c.JSON(ErrResp("id要是数字类型"))
		return
	}
	task, err := db.GetTaskById(taskIdNum)
	if err != nil {
		c.JSON(ErrResp(err.Error()))
		return
	}
	c.JSON(SuccessResp(task))
	return
}

func GetTaskList(c *gin.Context) {
	rp := &ReqPaginater{}
	err := c.ShouldBindJSON(&rp)
	if err != nil {
		c.JSON(ErrResp(err.Error()))
		return
	}
	taskList := db.GetTaskAll(rp.Page, rp.Size)
	c.JSON(SuccessResp(taskList))
	return
}
