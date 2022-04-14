/*
 * @Date: 2022-04-11 15:19:08
 * @LastEditors: recar
 * @LastEditTime: 2022-04-14 14:35:20
 */
package reverse

import (
	"bytes"
	"fmt"
	"net/url"
	"scanner/pkg/conf"
	"scanner/pkg/util"
	"time"
	"woodpecker/pkg/parse/cel/proto"

	"github.com/valyala/fasthttp"
)

// use ceye api

func NewReverse() *proto.Reverse {
	ceyeDomain := conf.GlobalConfig.Reverse.Domain
	flag := util.RandLowLetterNumber(8)
	if ceyeDomain == "" {
		return &proto.Reverse{}
	}
	urlStr := fmt.Sprintf("http://%s.%s", flag, ceyeDomain)
	u, _ := url.Parse(urlStr)
	return &proto.Reverse{
		Flag:               flag,
		Url:                util.ParseUrl(u),
		Domain:             u.Hostname(),
		Ip:                 "",
		IsDomainNameServer: false,
	}
}

func ReverseCheck(r *proto.Reverse, timeout int64) bool {
	ceyeApiToken := conf.GlobalConfig.Reverse.ApiKey
	if ceyeApiToken == "" || r.Domain == "" {
		return false
	}
	// 延迟 x 秒获取结果
	time.Sleep(time.Second * time.Duration(timeout))

	//check dns
	verifyUrl := fmt.Sprintf("http://api.ceye.io/v1/records?token=%s&type=dns&filter=%s", ceyeApiToken, r.Flag)
	if GetReverseResp(verifyUrl) {
		return true
	} else {
		//	check request
		verifyUrl := fmt.Sprintf("http://api.ceye.io/v1/records?token=%s&type=http&filter=%s", ceyeApiToken, r.Flag)
		if GetReverseResp(verifyUrl) {
			return true
		}
	}
	return false
}

func GetReverseResp(verifyUrl string) bool {
	notExist := []byte(`"data": []`)

	fastReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(fastReq)
	fastReq.SetRequestURI(verifyUrl)
	fastReq.Header.SetMethod(fasthttp.MethodGet)

	resp, err := util.DoFasthttpRequest(fastReq, false)

	if err != nil {
		return false
	}
	if !bytes.Contains(resp.Body, notExist) { // api返回结果不为空
		return true
	}
	return false
}
