/*
 * @Date: 2022-04-18 10:12:24
 * @LastEditors: recar
 * @LastEditTime: 2022-04-22 10:24:14
 */
/*
 * @Date: 2022-04-18 10:05:08
 * @LastEditors: recar
 * @LastEditTime: 2022-04-18 10:12:02
 */

package db

type Poc struct {
	Id      int `gorm:"primaryKey"`
	PocId   string
	PocName string
	PocDesc string
	Level   int
	Content string

	Enable        int
	Cve           string
	Source        string
	DescId        int
	RuleIds       string
	AppId         int
	FingerprintId int
	CreateTime    string
	UpdateTime    string
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

func AddPoc(poc *Poc) {
	currentTime := GetCurrentTime()
	poc.CreateTime = currentTime
	poc.UpdateTime = currentTime
	poc.Enable = 1
	SqlDb.Create(poc)
}

func DelPocById(id int) {
	SqlDb.Delete(Poc{}, id)
}

func UpdatePocBy(poc Poc) {
	poc.UpdateTime = GetCurrentTime()
	SqlDb.Save(poc)
}

func GetPocAll(page int, size int) []Poc {
	pocs := []Poc{}
	offset := (page - 1) * size
	SqlDb.Offset(offset).Limit(size).Find(&pocs)
	return pocs
}

func GetPocById(id int) (Poc, error) {
	poc := Poc{}
	result := SqlDb.Find(&poc, id)
	return poc, result.Error
}
