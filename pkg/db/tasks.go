/*
 * @Date: 2022-04-18 10:05:08
 * @LastEditors: recar
 * @LastEditTime: 2022-04-18 16:40:18
 */

package db

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id          uint `gorm:"primaryKey"`
	TaskName    string
	TaskDesc    string
	Target      string
	Port        string
	ScanType    int
	Status      int
	PocIds      string
	WhiteList   string
	TotalNumber int
	TestNumber  int
	CreateTime  time.Time
	UpdateTime  time.Time
}

func AddTask(task *Task) *gorm.DB {
	task.CreateTime = time.Now()
	task.UpdateTime = time.Now()
	result := SqlDb.Create(task)
	return result

}
