/*
 * @Date: 2022-04-11 14:49:57
 * @LastEditors: recar
 * @LastEditTime: 2022-04-22 17:35:45
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
	wapi := WebAPI.Group("/api")
	{
		// task
		wapi.POST("/task", AddTask)
		wapi.DELETE("/task", DelTask)
		wapi.POST("/start", StartTask)
		wapi.POST("/stop", StopTask)
		wapi.GET("/task", GetTaskList)
		wapi.GET("/task/:id", GetTaskInfo)
		//poc
		wapi.POST("/poc", AddPoc)
		wapi.DELETE("/poc", DelPoc)
		wapi.PUT("/poc", UpdatePoc)
		wapi.GET("/poc", GetPocList)
		wapi.GET("/poc/:id", GetPocInfo)
	}

	WebAPI.Run(listen)
}
