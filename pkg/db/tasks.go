/*
 * @Date: 2022-04-18 10:05:08
 * @LastEditors: recar
 * @LastEditTime: 2022-04-20 14:36:00
 */

package db

import (
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
	CreateTime  string
	UpdateTime  string
}

func AddTask(task *Task) *gorm.DB {
	currentTime := GetCurrentTime()
	task.CreateTime = currentTime
	task.UpdateTime = currentTime
	result := SqlDb.Create(task)
	return result

}

func GetTaskById(id int) (Task, error) {
	task := Task{}
	result := SqlDb.Find(&task, id)
	return task, result.Error
}

func GetTaskByName(name string) (Task, error) {
	task := Task{}
	result := SqlDb.Find(&task, name)
	return task, result.Error
}

func UpdateTaskStatusById(id int, status int) error {
	task, err := GetTaskById(id)
	currentTime := GetCurrentTime()
	task.UpdateTime = currentTime
	task.Status = status
	SqlDb.Save(task)
	return err
}

func DelTaskById(id int) {
	SqlDb.Delete(Task{}, id)
}

func GetTaskAll(page int, size int) []Task {
	tasks := []Task{}
	offset := (page - 1) * size
	SqlDb.Offset(offset).Limit(size).Find(&tasks)
	return tasks
}
