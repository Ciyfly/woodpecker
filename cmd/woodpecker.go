/*
 * @Date: 2022-04-13 16:44:15
 * @LastEditors: recar
 * @LastEditTime: 2022-04-19 10:52:18
 */
package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"woodpecker/api"
	"woodpecker/pkg/conf"
	"woodpecker/pkg/db"
	"woodpecker/pkg/log"
	"woodpecker/pkg/parse"
	"woodpecker/pkg/scan"

	"github.com/urfave/cli/v2"
)

func init() {
	fmt.Println(conf.Banner)
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt, os.Kill, syscall.SIGKILL)
	go func() {
		s := <-c
		log.Debug(fmt.Sprintf("recv signal: %d", s))
		log.Info("ctrl+c exit")
		os.Exit(0)
	}()
}

func main() {
	SetupCloseHandler()
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
		log.Errorf("cli.RunApp err: %s", err.Error())
		return
	}
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
	// init task result
	conf.Load()
	scan.InitTaskChannel()
	scan.InitResultChannel(mode)
	log.Init()

	// 统一转为 list 中 http://ip:port 的形式
	if mode == parse.ModeScan {
		// 处理target
		targets, err := parse.ParseTargets(target, port)
		if err != nil {
			log.Errorf(err.Error())
			os.Exit(1)
		}
		if targets == nil && mode == parse.ModeScan {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
		log.Debugf("ips: ", targets)
		// 通过生产者创建任务推到任务队列
		scan.ProducerTask(mode, targets, pocNames, nil, 0)
		end := <-scan.EndChannel
		log.Infof("Scan %s", end)
	} else {
		// web server
		db.InitDb()
		api.Server(server)
	}

	return nil
}
