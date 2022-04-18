/*
 * @Date: 2022-04-18 10:15:50
 * @LastEditors: recar
 * @LastEditTime: 2022-04-18 18:22:30
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
	"time"

	"gorm.io/gorm"
)

type Report struct {
	Id         uint `gorm:"primaryKey"`
	TaskId     int
	PocId      int
	Target     string
	Port       string
	Status     int
	Req        string
	Rsp        string
	CreateTime time.Time
	UpdateTime time.Time
}

func AddReport(report *Report) *gorm.DB {
	report.CreateTime = time.Now()
	report.UpdateTime = time.Now()
	result := SqlDb.Create(report)
	return result
}
