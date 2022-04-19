/*
 * @Date: 2022-04-18 10:12:24
 * @LastEditors: recar
 * @LastEditTime: 2022-04-19 10:00:41
 */
/*
 * @Date: 2022-04-18 10:05:08
 * @LastEditors: recar
 * @LastEditTime: 2022-04-18 10:12:02
 */

package db

import (
	"time"
)

type Poc struct {
	Id            int `gorm:"primaryKey"`
	PocId         string
	PocName       string
	Level         int
	Content       string
	Enable        bool
	Cve           string
	PocDesc       string
	Source        string
	DescId        int
	RuleIds       string
	Appid         int
	FingerpringId int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func GetPocByIds(pocIds []int) ([]Poc, error) {
	pocs := []Poc{}
	result := SqlDb.Find(&pocs, pocIds)
	return pocs, result.Error
}

func GetPocByName(name string) (Poc, error) {
	poc := Poc{}
	result := SqlDb.Where("poc_name = ?", name).First(&poc)
	return poc, result.Error
}
