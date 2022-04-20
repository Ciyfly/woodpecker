/*
 * @Date: 2022-04-18 10:15:50
 * @LastEditors: recar
 * @LastEditTime: 2022-04-20 11:05:24
 */
/*
 * @Date: 2022-04-18 10:12:24
 * @LastEditors: recar
 * @LastEditTime: 2022-04-18 10:13:38
 */
/*
 * @Date: 2022-04-18 10:05:08
 * @LastEditors: recar
 * @LastEditTime: 2022-04-18 10:12:02
 */

package db

import (
	"gorm.io/gorm"
)

type Report struct {
	Id         int `gorm:"primaryKey"`
	TaskId     int
	PocId      int
	Target     string
	Port       string
	Status     int
	Req        string
	Rsp        string
	CreateTime string
	UpdateTime string
}

func AddReport(report *Report) *gorm.DB {
	result := SqlDb.Create(report)
	return result
}
