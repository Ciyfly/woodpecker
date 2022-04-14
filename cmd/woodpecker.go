/*
 * @Date: 2022-04-13 16:44:15
 * @LastEditors: recar
 * @LastEditTime: 2022-04-14 14:20:40
 */
package main

import (
	"net/url"
	"os"
	"strings"
	"woodpecker/api"
	"woodpecker/pkg/conf"
	"woodpecker/pkg/log"
	"woodpecker/pkg/scan"
	"woodpecker/pkg/util"

	"github.com/urfave/cli/v2"
)

func init() {
	// 初始化 Task 任务队列
	scan.InitTaskChannel()
}

func main() {
	app := cli.NewApp()
	app.Name = conf.ServiceName
	app.Usage = conf.Website
	app.Version = conf.Version

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			// 模式 server scan
			Name:    "mode",
			Aliases: []string{"m"},
			Value:   "scan",
			Usage:   "启动模式 scan是命令行扫描 server是后端模式",
		},
		&cli.StringFlag{
			// 服务端地址
			Name:    "server",
			Aliases: []string{"s"},
			Value:   "127.0.0.1:8787",
			Usage:   "后端监听地址",
		},
		&cli.StringFlag{
			// 指定poc name
			Name:    "poc_names",
			Aliases: []string{"pn"},
			Usage:   "指定poc 列表 逗号隔开",
		},
		&cli.StringFlag{
			// 目标地址
			Name:    "target",
			Aliases: []string{"t"},
			Value:   "",
			Usage:   "扫描的目标 支持网段 ip 网址",
		},
		&cli.StringFlag{
			// 目标端口
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "",
			Usage:   "当使用网段和ip的时候需要指定扫描的端口",
		},
	}
	app.Action = RunMain

	err := app.Run(os.Args)
	if err != nil {
		log.Errorf("cli.RunApp err: %v", err)
		return
	}
}

func handlerTargets(target string, port string) []string {
	targets := []string{}
	if target == "" {
		log.Info("扫描需要输入目标地址")
		return nil
	}
	if strings.Index(target, "http") != -1 {
		u, err := url.Parse(target)
		if err != nil {
			log.Info("url 不符合规范  http://127.0.0.1")
			return nil
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
			log.Info("使用网段的话需要输入端口")
			return nil
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
				log.Error("使用ip的话需要输入端口")
				return nil
			}
			targets = append(targets, "http://"+target+":"+port)
		}

	}
	return targets
}

func RunMain(c *cli.Context) error {
	var pocNames []string

	mode := c.String("mode")
	server := c.String("server")
	target := c.String("target")
	port := c.String("port")
	cmdPocNames := c.String("poc_names")
	if cmdPocNames != "" {
		pocNames = strings.Split(cmdPocNames, ",")
	}
	// 统一转为 list 中 http://ip:port 的形式
	if mode == "scan" {
		// 处理target
		targets := handlerTargets(target, port)
		if targets == nil && mode == "scan" {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
		log.Debugf("ips: ", targets)
		// 通过生产者创建任务推到任务队列
		scan.ProducerTask(mode, targets, pocNames, nil)
		end := <-scan.EndChannel
		log.Infof("Scan %s", end)
	} else {
		// web server
		api.Server(server)
	}

	return nil
}
