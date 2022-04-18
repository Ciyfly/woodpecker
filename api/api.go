/*
 * @Date: 2022-04-11 14:49:57
 * @LastEditors: recar
 * @LastEditTime: 2022-04-18 15:27:14
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
	WebAPI.POST("/task", AddTask)

	WebAPI.Run(listen)
}
