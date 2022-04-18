/*
 * @Date: 2022-04-18 17:51:34
 * @LastEditors: recar
 * @LastEditTime: 2022-04-18 17:54:29
 */
package parse

func ParseTargets(target string, port string) ([]string, error) {
	targets := []string{}
	if target == "" {
		return targets, errors.New("扫描需要输入目标地址")
	}
	if strings.Index(target, "http") != -1 {
		u, err := url.Parse(target)
		if err != nil {
			return targets, errors.New("url 不符合规范  http://127.0.0.1")
		}
		// http://127.0.0.1:8080
		if strings.LastIndex(target, ":") != 4 {
			targets = append(targets, "http://"+u.Host)
		} else {
			if strings.Index(target, "https") != -1 {
				targets = append(targets, "https://"+u.Host+":"+"443")
			} else {
				targets = append(targets, "http://"+u.Host+":"+"80")
			}
		}
	} else if strings.Index(target, "http") == -1 && strings.Index(target, "/") != -1 {
		// 网段
		if port == "" {
			return targets, errors.New("使用网段的话需要输入端口")
		}
		ips := util.Ips2List(target)
		for _, ip := range ips {
			targets = append(targets, "http://"+ip+":"+port)
		}
	} else {
		// 127.0.0.1:80
		if strings.Index(target, ":") != -1 {
			targets = append(targets, target)
		} else {
			// -t 127.0.0.1 -p 80
			if port == "" {
				return targets, errors.New("使用ip的话需要输入端口")
			}
			targets = append(targets, "http://"+target+":"+port)
		}

	}
	return targets, nil
}
