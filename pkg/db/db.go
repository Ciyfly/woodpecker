/*
 * @Date: 2022-04-18 10:18:23
 * @LastEditors: recar
 * @LastEditTime: 2022-04-20 11:15:21
 */
package db

import (
	"os"
	"path"
	"path/filepath"
	"time"
	"woodpecker/pkg/log"

	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
)

var SqlDb *gorm.DB

func InitDb() {
	baseDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Errorf("InitDb get base dir error: %s", err.Error())
	}
	dbPath := path.Join(baseDir, "woodpecker.db")
	// github.com/mattn/go-sqlite3
	SqlDb, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Errorf("InitDb error: %s", err.Error())
	}
	// SqlDb.SingularTable(true) 可以取消表名的复数形式，使得表名和结构体名称一致
	SqlDb.AutoMigrate(&App{}, &Poc{}, &Report{}, &Task{})
}

func GetCurrentTime() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime
}
