/*
 * @Date: 2022-04-11 14:49:57
 * @LastEditors: recar
 * @LastEditTime: 2022-04-20 14:46:42
 */
package api

import (
	"woodpecker/pkg/log"

	"github.com/gin-gonic/gin"
)

var WebAPI *gin.Engine

func Server(listen string) {
	gin.SetMode(gin.ReleaseMode)
	WebAPI = gin.Default()
	log.Infof("server start:  %s", listen)
	// task
	taskApi := WebAPI.Group("/task")
	{
		taskApi.POST("/addtask", AddTask)
		taskApi.POST("/deltask", DelTask)
		taskApi.POST("/start", StartTask)
		taskApi.POST("/stop", StopTask)
		taskApi.GET("/gettasklist", GetTaskList)
		taskApi.POST("/gettaskinfo", GetTaskInfo)
	}
	// poc
	pocApi := WebAPI.Group("/poc")
	{
		pocApi.POST("/add", AddPoc)
		pocApi.POST("/del", DelPoc)
		pocApi.POST("/modify", UpdatePoc)
		pocApi.POST("/getpoclist", GetPocList)
		pocApi.POST("/getpocinfo", GetPocInfo)
	}

	WebAPI.Run(listen)
}
