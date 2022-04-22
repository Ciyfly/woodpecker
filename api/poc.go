/*
 * @Date: 2022-04-20 10:05:31
 * @LastEditors: recar
 * @LastEditTime: 2022-04-22 17:42:09
 */
package api

import (
	"strconv"
	"woodpecker/pkg/db"

	"github.com/gin-gonic/gin"
)

type ReqPoc struct {
	Id      int
	PocId   string `json:"poc_id"`
	PocName string `json:"poc_name"`
	PocDesc string `json:"poc_desc"`
	RuleId  string `json:"rule_id"`

	Content       string `json:"content,omitempty"`
	FingerprintId int    `json:"fingerprint,omitempty"`
	AppId         int    `json:"app_id,omitempty"`
	DescId        int    `json:"desc_id,omitempty"`
	Cve           string `json:"cve,omitempty"`
	Source        string `json:"source,omitempty"`
	Level         int    `json:"level,omitempty"`
}

func AddPoc(c *gin.Context) {
	reqPoc := ReqPoc{}
	err := c.ShouldBindJSON(&reqPoc)
	if err != nil {
		c.JSON(ErrResp(err.Error()))
		return
	}
	_, err = db.GetPocByName(reqPoc.PocName)
	if err == nil {
		c.JSON(ErrResp("poc name 已经存在"))
		return
	}
	// 后续可能更多操作 例如判断规则id是否存在 等等
	poc := &db.Poc{
		PocId:         reqPoc.PocId,
		PocName:       reqPoc.PocName,
		PocDesc:       reqPoc.PocDesc,
		Level:         reqPoc.Level,
		Content:       reqPoc.Content,
		Cve:           reqPoc.Cve,
		Source:        reqPoc.Source,
		DescId:        reqPoc.DescId,
		RuleIds:       reqPoc.RuleId,
		AppId:         reqPoc.AppId,
		FingerprintId: reqPoc.FingerprintId,
	}
	db.AddPoc(poc)
	c.JSON(SuccessResp(poc))
	return
}

func DelPoc(c *gin.Context) {
	reqPoc := ReqPoc{}
	err := c.ShouldBindJSON(&reqPoc)
	if err != nil {
		c.JSON(ErrResp(err.Error()))
		return
	}
	db.DelPocById(reqPoc.Id)
	c.JSON(SuccessResp(""))
	return
}

func UpdatePoc(c *gin.Context) {
	reqPoc := ReqPoc{}
	err := c.ShouldBindJSON(&reqPoc)
	if err != nil {
		c.JSON(ErrResp(err.Error()))
		return
	}
	dbPoc, err := db.GetPocByName(reqPoc.PocName)
	if err != nil {
		c.JSON(ErrResp(err.Error()))
		return
	}
	// 后续可能更多操作 例如判断规则id是否存在 等等
	dbPoc.PocId = reqPoc.PocId
	dbPoc.PocName = reqPoc.PocName
	dbPoc.PocDesc = reqPoc.PocDesc
	dbPoc.Content = reqPoc.Content
	dbPoc.Cve = reqPoc.Cve
	dbPoc.Source = reqPoc.Source
	dbPoc.DescId = reqPoc.DescId
	dbPoc.RuleIds = reqPoc.RuleId
	dbPoc.AppId = reqPoc.AppId
	dbPoc.FingerprintId = reqPoc.FingerprintId
	db.UpdatePocBy(dbPoc)
	c.JSON(SuccessResp(dbPoc))
	return
}

func GetPocList(c *gin.Context) {
	rp := &ReqPaginater{}
	err := c.ShouldBindJSON(&rp)
	if err != nil {
		c.JSON(ErrResp(err.Error()))
		return
	}
	pocList := db.GetPocAll(rp.Page, rp.Size)
	c.JSON(SuccessResp(pocList))
	return
}

func GetPocInfo(c *gin.Context) {
	pocId := c.Param("id")
	pocIdNum, err := strconv.Atoi(pocId)
	if err != nil {
		c.JSON(ErrResp("id要是数字类型"))
		return
	}
	poc, err := db.GetPocById(pocIdNum)
	if err != nil {
		c.JSON(ErrResp(err.Error()))
		return
	}
	c.JSON(SuccessResp(poc))
	return
}
