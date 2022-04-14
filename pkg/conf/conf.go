/*
 * @Date: 2022-04-13 16:48:45
 * @LastEditors: recar
 * @LastEditTime: 2022-04-14 18:50:13
 */
package conf

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	Banner = `
	██╗    ██╗ ██████╗  ██████╗ ██████╗ ██████╗ ███████╗ ██████╗██╗  ██╗███████╗██████╗ 
	██║    ██║██╔═══██╗██╔═══██╗██╔══██╗██╔══██╗██╔════╝██╔════╝██║ ██╔╝██╔════╝██╔══██╗
	██║ █╗ ██║██║   ██║██║   ██║██║  ██║██████╔╝█████╗  ██║     █████╔╝ █████╗  ██████╔╝
	██║███╗██║██║   ██║██║   ██║██║  ██║██╔═══╝ ██╔══╝  ██║     ██╔═██╗ ██╔══╝  ██╔══██╗
	╚███╔███╔╝╚██████╔╝╚██████╔╝██████╔╝██║     ███████╗╚██████╗██║  ██╗███████╗██║  ██║
	 ╚══╝╚══╝  ╚═════╝  ╚═════╝ ╚═════╝ ╚═╝     ╚══════╝ ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
																						
	`
	ServiceName = "woodpecker"
	Version     = "1.0"
	Website     = ""
	// 限流控制并发
	RateSize = 10
	// 超时时间
	ConnectTimeOut = 3
	RecvTimeOut    = 5
)

var GlobalConfig *Conf

type Reverse struct {
	ApiKey string `mapstructure:"api_key"`
	Domain string `mapstructure:"domain"`
}

type LogConfig struct {
	MaxSize    int  `mapstructure:"max_size"`
	MaxBackups int  `mapstructure:"max_backups"`
	MaxAge     int  `mapstructure:"max_age"`
	Compress   bool `mapstructure:"compress"`
}

type HttpConfig struct {
	ConnectTimeOut int `mapstructure:"connect_timeout"`
	RecvTimeOut    int `mapstructure:"recv_timeout"`
}

type Conf struct {
	LogConfig  LogConfig  `mapstructure:"logConfig"`
	Reverse    Reverse    `mapstructure:"reverse"`
	HttpConfig HttpConfig `mapstructure:"httpConfig"`
	RateSize   int        `mapstructure:"rate"`
}

func ReadYamlConfig(configFile string) {
	// 加载config
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("load config, fail to read 'config.yaml': %v\n", err)
	}
	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		fmt.Printf("load config, fail to parse 'config.yaml', check format: %v\n", err)
	}
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Load() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("load config error: %v\n", err)
	}
	configFile := path.Join(dir, "configs", "config.yml")
	if !Exists(configFile) {
		fmt.Println("配置文件不存在")
	}
	ReadYamlConfig(configFile)
}
