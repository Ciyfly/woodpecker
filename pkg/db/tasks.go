/*
 * @Date: 2022-04-18 10:05:08
 * @LastEditors: recar
 * @LastEditTime: 2022-04-22 15:30:20
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

func GetTaskTotalNumberById(id int) (int, error) {
	task, err := GetTaskById(id)
	return task.TotalNumber, err
}

func GetTaskTestNumberById(id int) (int, error) {
	task, err := GetTaskById(id)
	return task.TestNumber, err
}

func UpdateTaskStatusById(id int, status int) {
	currentTime := GetCurrentTime()
	SqlDb.Model(&Task{}).Where("id = ?", id).Updates(map[string]interface{}{"status": status, "update_time": currentTime})
}

func UpdateTaskTotalNumberById(id int, totalNumber int) {
	currentTime := GetCurrentTime()
	SqlDb.Model(&Task{}).Where("id = ?", id).Updates(map[string]interface{}{"total_number": totalNumber, "update_time": currentTime})
}

func UpdateTaskTestNumberById(id int, testNumber int) {
	currentTime := GetCurrentTime()
	SqlDb.Model(&Task{}).Where("id = ?", id).Updates(map[string]interface{}{"test_number": testNumber, "update_time": currentTime})
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
