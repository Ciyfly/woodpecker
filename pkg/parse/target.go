/*
 * @Date: 2022-04-18 17:51:34
 * @LastEditors: recar
 * @LastEditTime: 2022-04-19 10:10:57
 */
package parse

import (
	"errors"
	"net/url"
	"strings"
	"woodpecker/pkg/util"
)

func ParseTargets(target string, ports string) ([]string, error) {
	targets := []string{}
	if target == "" {
		return targets, errors.New("扫描需要输入目标地址")
	}
	if strings.Index(target, "http") != -1 {
		u, err := url.Parse(target)
		if err != nil {
			return targets, errors.New("url 不符合规范  http://127.0.0.1")
		}
		baseTarget := u.Scheme + "://" + u.Host
		targets = append(targets, baseTarget)
	} else if strings.Index(target, "http") == -1 && strings.Index(target, "/") != -1 {
		// 网段
		if ports == "" {
			return targets, errors.New("使用网段的话需要输入端口")
		}
		ips := util.Ips2List(target)
		for _, ip := range ips {
			for _, port := range strings.Split(ports, ",") {
				targets = append(targets, "http://"+ip+":"+port)
			}
		}
	} else {
		// 127.0.0.1:80
		if strings.Index(target, ":") != -1 {
			targets = append(targets, target)
		} else {
			// -t 127.0.0.1 -p 80
			if ports == "" {
				return targets, errors.New("使用ip的话需要输入端口")
			}
			for _, port := range strings.Split(ports, ",") {
				targets = append(targets, "http://"+target+":"+port)
			}
		}

	}
	return targets, nil
}
