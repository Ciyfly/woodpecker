/*
 * @Date: 2022-04-14 16:24:58
 * @LastEditors: recar
 * @LastEditTime: 2022-04-15 18:01:38
 */

package parse

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
	"woodpecker/pkg/cel/proto"
	"woodpecker/pkg/conf"
	"woodpecker/pkg/log"
)

func UrlToPUrl(url *url.URL) *proto.UrlType {
	nu := &proto.UrlType{}
	nu.Scheme = url.Scheme
	nu.Domain = url.Hostname()
	nu.Host = url.Host
	nu.Port = url.Port()
	nu.Path = url.EscapedPath()
	nu.Query = url.RawQuery
	nu.Fragment = url.Fragment
	return nu
}

func GetPReqByTarget(target string) *proto.Request {
	pReq := &proto.Request{}
	originalReq, err := http.NewRequest("GET", target, nil)
	if err != nil {
		log.Errorf("target: %s req error: %s", target, err.Error())
	}
	pReq.Method = originalReq.Method
	pReq.Url = UrlToPUrl(originalReq.URL)
	// headers := make(map[string]string)
	// for k := range originalReq.Header {
	// 	headers[k] = originalReq.Header.Get(k)
	// }
	return pReq
}

func DoRequest(rule Rule, target string) (*proto.Response, error) {
	// 解析url 获取base path
	parseUrl, err := url.Parse(target)
	if err != nil {
		log.Errorf("parse url: %s error: %s", target, err.Error())
	}
	TargetBaseUrl := parseUrl.Scheme + "://" + parseUrl.Host
	targetUrl := TargetBaseUrl + rule.Request.Path
	targetUrl = strings.ReplaceAll(targetUrl, " ", "%20")
	targetUrl = strings.ReplaceAll(targetUrl, "+", "%20")
	body := bytes.NewBuffer([]byte(rule.Request.Body))
	log.Debugf("method: %s", rule.Request.Method)
	log.Debugf("targetUrl: %s", targetUrl)
	log.Debugf("body: %s", body)
	req, err := http.NewRequest(rule.Request.Method, targetUrl, body)
	if err != nil {
		log.Errorf("Error creating request: %s", err.Error())
	}
	for key, value := range rule.Request.Headers {
		req.Header.Set(key, value)
	}

	// create http client
	client := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(conf.GlobalConfig.HttpConfig.ConnectTimeOut)) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(time.Second * time.Duration(conf.GlobalConfig.HttpConfig.RecvTimeOut))) //设置发送接收数据超时
				return c, nil
			},
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	if rule.Request.FollowRedirects {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return fmt.Errorf("Not Redirect")
		}
	}
	// cont-type
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	rsp, err := client.Do(req)
	// rsp 转proto Response
	if err != nil {
		return nil, err
	}
	protoResponse := &proto.Response{}
	headers := make(map[string]string)
	for k := range rsp.Header {
		headers[k] = strings.ToLower(rsp.Header.Get(k))
	}
	protoResponse.Url = UrlToPUrl(parseUrl)
	protoResponse.Status = int32(rsp.StatusCode)
	protoResponse.Headers = headers
	protoResponse.ContentType = rsp.Header.Get("Content-Type")
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Errorf("url: %s  parse rsp body error: %s", targetUrl, err.Error())
	}
	protoResponse.Body = rspBody
	return protoResponse, err
}
