/*
 * @Date: 2022-04-13 16:55:00
 * @LastEditors: recar
 * @LastEditTime: 2022-04-29 18:14:21
 */
package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"woodpecker/pkg/cel/proto"
	"woodpecker/pkg/log"

	"net"
	"os"
	"path"
)

func increment(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}

func Ips2List(ips string) []string {
	targets := []string{}
	ipAddr, ipNet, err := net.ParseCIDR(ips)
	if err != nil {
		log.Errorf("网段格式不对 请检查: %s", ips)
		os.Exit(0)
	}

	for ip := ipAddr.Mask(ipNet.Mask); ipNet.Contains(ip); increment(ip) {
		targets = append(targets, ip.String())
	}

	// 去除第一个ip 因为是网段的
	if len(targets) >= 2 {
		targets = targets[1:]
	}
	return targets
}

func GetDirFilePath(dir string) []string {
	filePathList := []string{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Errorf("获取目录下文件报错: %s %s", dir, err.Error())
	}

	for _, file := range files {
		filePathList = append(filePathList, path.Join(dir, file.Name()))
	}
	return filePathList
}

func GetAllFile(dirPath string) []string {
	allFilePaths := []string{}
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		allFilePaths = append(allFilePaths, path)
		return nil
	})
	if err != nil {
		log.Errorf("GetAllFile error: %s", err.Error())
	}
	return allFilePaths
}

func ValInSliceStr(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func FileIsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func UrlTypeToString(u *proto.UrlType) string {
	var buf strings.Builder
	if u.Scheme != "" {
		buf.WriteString(u.Scheme)
		buf.WriteByte(':')
	}
	if u.Scheme != "" || u.Host != "" {
		if u.Host != "" || u.Path != "" {
			buf.WriteString("//")
		}
		if h := u.Host; h != "" {
			buf.WriteString(u.Host)
		}
	}
	path := u.Path
	if path != "" && path[0] != '/' && u.Host != "" {
		buf.WriteByte('/')
	}
	if buf.Len() == 0 {
		if i := strings.IndexByte(path, ':'); i > -1 && strings.IndexByte(path[:i], '/') == -1 {
			buf.WriteString("./")
		}
	}
	buf.WriteString(path)

	if u.Query != "" {
		buf.WriteByte('?')
		buf.WriteString(u.Query)
	}
	if u.Fragment != "" {
		buf.WriteByte('#')
		buf.WriteString(u.Fragment)
	}
	return buf.String()
}

func Req2Text(req *http.Request) string {
	var text string
	// 补全url path params
	if req == nil {
		return text
	}
	if req.URL.RawQuery == "" {
		text = fmt.Sprintf("%s %s %s\r\n", req.Method, req.URL.Path, req.Proto)
	} else {
		text = fmt.Sprintf("%s %s?%s %s\r\n", req.Method, req.URL.Path, req.URL.RawQuery, req.Proto)
	}
	for k, _ := range req.Header {
		text += k + ":" + req.Header.Get(k) + "\r\n"
	}
	// 补全host request会默认移除掉host
	text += "Host:" + req.Host + "\r\n"
	// 补全UA
	if req.Header.Get("user-agent") == "" {
		text += "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36" + "\r\n"
	}
	text += "\r\n"
	reqBody, err := ioutil.ReadAll(req.Body)
	if err == nil {
		text += string(reqBody)
	}
	// log.Errorf("req: %s", text)
	return text
}

func Rsp2Text(rsp *http.Response) string {
	var text string
	if rsp == nil {
		return text
	}
	text = fmt.Sprintf("%s %s\r\n", rsp.Proto, rsp.Status)
	for k, _ := range rsp.Header {
		text += k + ":" + rsp.Header.Get(k) + "\r\n"
	}
	text += "\r\n"
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err == nil {
		text += string(rspBody)
	}
	// log.Errorf("rsp: %s", text)
	return text
}
