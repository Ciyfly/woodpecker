/*
 * @Date: 2022-04-18 10:05:08
 * @LastEditors: recar
 * @LastEditTime: 2022-04-19 10:04:40
 */

package db

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id          int `gorm:"primaryKey"`
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
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func AddTask(task *Task) *gorm.DB {
	result := SqlDb.Create(task)
	return result

}

func GetTaskById(id int) (Task, error) {
	task := Task{}
	result := SqlDb.Find(&task, id)
	return task, result.Error
}

func UpdateTaskStatusById(id int, status int) error {
	task, err := GetTaskById(id)
	task.Status = status
	SqlDb.Save(task)
	return err
}
