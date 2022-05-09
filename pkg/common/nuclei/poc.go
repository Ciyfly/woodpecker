/*
 * @Date: 2022-04-29 15:05:21
 * @LastEditors: recar
 * @LastEditTime: 2022-04-29 19:25:14
 */
package nuclei

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"woodpecker/pkg/db"
	"woodpecker/pkg/log"
	"woodpecker/pkg/util"

	"go.uber.org/ratelimit"

	"github.com/projectdiscovery/nuclei/v2/pkg/catalog"
	"github.com/projectdiscovery/nuclei/v2/pkg/output"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/protocolinit"
	"github.com/projectdiscovery/nuclei/v2/pkg/templates"
	"github.com/projectdiscovery/nuclei/v2/pkg/types"
)

type Poc = templates.Template

var ExecuterOptions protocols.ExecuterOptions

func InitNucleiExecuterOptions(limiter int, timeout int) error {
	fakeWriter := fakeWrite{}
	progress := &fakeProgress{}
	o := types.Options{
		RateLimit:               limiter,
		BulkSize:                25,
		TemplateThreads:         25,
		HeadlessBulkSize:        10,
		HeadlessTemplateThreads: 10,
		Timeout:                 timeout,
		Retries:                 1,
		MaxHostError:            30,
	}
	err := protocolinit.Init(&o)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not initialize protocols: %s", err))
	}
	catalog2 := catalog.New("")
	ExecuterOptions = protocols.ExecuterOptions{
		Output:      &fakeWriter,
		Options:     &o,
		Progress:    progress,
		Catalog:     catalog2,
		RateLimiter: ratelimit.New(limiter),
	}
	return nil
}

func ParsePocFile(filePath string) (*templates.Template, error) {
	var err error
	template, err := templates.Parse(filePath, nil, ExecuterOptions)
	if err != nil {
		return nil, err
	}
	if template == nil {
		return nil, nil
	}
	return template, nil
}

func ExecuteNucleiPoc(input string, poc *templates.Template) ([]string, error) {
	var ret []string
	var results bool
	e := poc.Executer
	name := fmt.Sprint(poc.ID)
	err := e.ExecuteWithResults(input, func(result *output.InternalWrappedEvent) {
		for _, r := range result.Results {
			results = true
			if r.ExtractorName != "" {
				ret = append(ret, name+":"+r.ExtractorName)
			} else if r.MatcherName != "" {
				ret = append(ret, name+":"+r.MatcherName)
			}
		}
	})
	if err != nil || !results {
		return nil, nil
	}
	if len(ret) == 0 {
		ret = append(ret, name)
	}
	return ret, err
}

func LoadNucleiYamlPoc(pocNames []string) []*Poc {
	// TODO: 对于多层文件poc name 拼接路径的问题处理优化
	baseDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	baseDir = strings.Replace(baseDir, "\\", "/", -1)
	if err != nil {
		log.Errorf("Nuclei Error loading YAML file: %s", err.Error())
	}
	pocDIr := path.Join(baseDir, "pocs", "nuclei")
	pocList := []*Poc{}
	allNucleiPocPathList := util.GetAllFile(pocDIr)
	for _, pocPath := range allNucleiPocPathList {
		// 过滤掉workflows nuclei 的workflows会默认用相对路径来找对于tag的poc
		if strings.Contains(pocPath, "workflows") {
			continue
		}
		fileName := filepath.Base(pocPath)
		// 如果不是yaml后缀的就不处理
		if find := strings.Contains(fileName, "yaml"); !find {
			continue
		}
		fileName = strings.Replace(fileName, ".yaml", "", -1)
		if len(pocNames) == 0 {
			poc, err := ParsePocFile(pocPath)
			if err != nil {
				log.Debugf("nuclei parse poc: %s error: %s", pocPath, err.Error())
			} else {
				pocList = append(pocList, poc)
			}
		} else {
			if util.ValInSliceStr(pocNames, fileName) == true {
				poc, err := ParsePocFile(pocPath)
				if err != nil {
					log.Debugf("nuclei parse poc: %s error: %s", pocPath, err.Error())
				} else {
					pocList = append(pocList, poc)
				}
			}
		}

	}
	return pocList
}

// db poc to yaml poc
func DbPoc2NucleiPoc(dbPocs []db.Poc) []*Poc {
	var nucleiPocs []*Poc
	for _, dp := range dbPocs {
		if dp.Source != "nuclei" {
			continue
		}
		// fork了nuclei 增加了一个解析conent内容的函数
		poc, err := templates.ParseContent([]byte(dp.Content), nil, ExecuterOptions)
		if err != nil {
			log.Errorf("nuclei parse poc error: %s", err.Error())
		}
		nucleiPocs = append(nucleiPocs, poc)
	}
	return nucleiPocs
}
