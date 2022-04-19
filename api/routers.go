/*
 * @Date: 2022-04-18 11:42:25
 * @LastEditors: recar
 * @LastEditTime: 2022-04-19 15:01:31
 */
package api

import (
	"fmt"
	"net/http"
	"strings"
	"woodpecker/pkg/db"
	"woodpecker/pkg/parse"
	"woodpecker/pkg/scan"

	"github.com/gin-gonic/gin"
)

const (
	SuccessCode = 1
	ErrCode     = 0
)

// API Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func SuccessResp(data interface{}) (int, Response) {
	res := Response{
		Code:  SuccessCode,
		Data:  data,
		Error: "",
	}
	return http.StatusOK, res
}

func ErrResp(errStr string) (int, Response) {
	res := Response{
		Code:  ErrCode,
		Data:  nil,
		Error: errStr,
	}
	return http.StatusOK, res
}

type ScanTask struct {
	PocIds   []int  `json:"poc_ids"`
	Port     string `json:"port"`
	Targets  string `json:"targets"`
	TaskDesc string `json:"task_desc"`
	TaskName string `json:"task_name"`
}

// router

func AddTask(c *gin.Context) {
	scanTask := ScanTask{}
	err := c.ShouldBindJSON(&scanTask)
	if err != nil {
		c.JSON(ErrResp(err.Error()))
		return
	}
	task := &db.Task{
		TaskName: scanTask.TaskName,
		TaskDesc: scanTask.TaskDesc,
		Target:   scanTask.Targets,
		Port:     scanTask.Port,
		PocIds:   strings.Trim(strings.Join(strings.Fields(fmt.Sprint(scanTask.PocIds)), ","), "[]"),
	}
	result := db.AddTask(task)
	if result.Error != nil {
		c.JSON(ErrResp(result.Error.Error()))
		return
	} else {
		targets, err := parse.ParseTargets(scanTask.Targets, scanTask.Port)
		if err != nil {
			c.JSON(ErrResp(err.Error()))
			return
		}
		scan.ProducerTask(parse.ModeServer, targets, nil, scanTask.PocIds, task.Id)
		c.JSON(SuccessResp(task.Id))
		return
	}
}
