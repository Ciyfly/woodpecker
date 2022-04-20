/*
 * @Date: 2022-04-18 11:42:25
 * @LastEditors: recar
 * @LastEditTime: 2022-04-20 14:53:36
 */
package api

import (
	"net/http"
)

const (
	SuccessCode = 200
	ErrorCode   = 500
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ReqPaginater struct {
	Page int
	Size int
}

func SuccessResp(data interface{}) (int, Response) {
	res := Response{
		Code: SuccessCode,
		Data: data,
	}
	return http.StatusOK, res
}

func ErrResp(errStr string) (int, Response) {
	res := Response{
		Code:    ErrorCode,
		Data:    nil,
		Message: errStr,
	}
	return http.StatusInternalServerError, res
}
