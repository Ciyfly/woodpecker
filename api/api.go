/*
 * @Date: 2022-04-11 14:49:57
 * @LastEditors: recar
 * @LastEditTime: 2022-04-13 16:58:21
 */
package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Server(listen string) {
	gin.SetMode(gin.ReleaseMode)
	web := gin.Default()
	log.Println("server start:", listen)
	web.Run(listen)
}
