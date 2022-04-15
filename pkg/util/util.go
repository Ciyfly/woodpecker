/*
 * @Date: 2022-04-13 16:55:00
 * @LastEditors: recar
 * @LastEditTime: 2022-04-15 17:38:41
 */
package util

import (
	"io/ioutil"
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
