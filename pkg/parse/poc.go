/*
 * @Date: 2022-04-13 17:17:01
 * @LastEditors: recar
 * @LastEditTime: 2022-04-14 19:21:17
 */
package parse

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"woodpecker/pkg/cel/proto"
	"woodpecker/pkg/log"
	"woodpecker/pkg/util"

	"gopkg.in/yaml.v2"
)

type RuleRequest struct {
	Cache           bool              `yaml:"cache"`
	Method          string            `yaml:"method"`
	Path            string            `yaml:"path"`
	Headers         map[string]string `yaml:"headers"`
	Body            string            `yaml:"body"`
	FollowRedirects bool              `yaml:"follow_redirects"`
	Content         string            `yaml:"content"`
	ReadTimeout     string            `yaml:"read_timeout"`
	ConnectionID    string            `yaml:"connection_id"`
}

type Output struct {
	Search string `yaml:"search"`
	Home   string `yaml:"home"`
}
type Rule struct {
	Request    RuleRequest `yaml:"request"`
	Expression string      `yaml:"expression"`
	Output     Output      `yaml:"output"`
}

type Detail struct {
	Author      string   `yaml:"author"`
	Links       []string `yaml:"links"`
	Description string   `yaml:"description"`
}

type Poc struct {
	Name       string            `yaml:"name"`
	Transport  string            `yaml:"transport"`
	Set        yaml.MapSlice     `yaml:"set"`
	Rules      map[string]Rule   `json:"rules"`
	Groups     map[string][]Rule `json:"groups"`
	Expression string            `yaml:"expression"`
	Detail     Detail            `yaml:"detail"`
}

func LoadYamlPoc(pocNames []string) []*Poc {
	baseDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	baseDir = strings.Replace(baseDir, "\\", "/", -1)
	if err != nil {
		log.Errorf("Error loading YAML file: %s", err.Error())
	}
	pocDIr := path.Join(baseDir, "pocs")
	pocList := []*Poc{}
	pocPathList := []string{}
	log.Debugf("pocNames: %v", pocNames)
	log.Debugf("pocNames: %d", len(pocNames))
	if len(pocNames) != 0 {
		for _, name := range pocNames {
			pocPathList = append(pocPathList, path.Join(pocDIr, name+".yml"))
		}
	} else {
		pocPathList = util.GetDirFilePath(pocDIr)
	}
	for _, pocPath := range pocPathList {
		log.Debugf("pocPath: %s", pocPath)
		poc := &Poc{}
		yamlFile, err := ioutil.ReadFile(pocPath)
		if err != nil {
			log.Errorf("Failed to read poc file %s %s", pocPath, err.Error())
		}
		err = yaml.Unmarshal(yamlFile, poc)
		if err != nil {
			log.Errorf("Failed to unmarshal poc file %s %s", pocPath, err.Error())
		}
		pocList = append(pocList, poc)
	}
	return pocList
}

func (rule *Rule) ReplaceSet(varMap map[string]interface{}) {
	for setKey, setValue := range varMap {
		// 过滤掉 map
		_, isMap := setValue.(map[string]string)
		if isMap {
			continue
		}
		value := fmt.Sprintf("%v", setValue)
		// 替换请求头中的 自定义字段
		for headerKey, headerValue := range rule.Request.Headers {
			rule.Request.Headers[headerKey] = strings.ReplaceAll(headerValue, "{{"+setKey+"}}", value)
		}
		// 替换请求路径中的 自定义字段
		rule.Request.Path = strings.ReplaceAll(strings.TrimSpace(rule.Request.Path), "{{"+setKey+"}}", value)
		// 替换body的 自定义字段
		rule.Request.Body = strings.ReplaceAll(strings.TrimSpace(rule.Request.Body), "{{"+setKey+"}}", value)
	}
}

// search
func (rule *Rule) ReplaceSearch(resp *proto.Response, varMap map[string]interface{}) map[string]interface{} {
	result := doSearch(strings.TrimSpace(rule.Output.Search), string(resp.Body))
	if result != nil && len(result) > 0 { // 正则匹配成功
		for k, v := range result {
			varMap[k] = v
		}
	}
	return varMap
}

// 实现 search
func doSearch(re string, body string) map[string]string {
	r, err := regexp.Compile(re)
	if err != nil {
		return nil
	}
	result := r.FindStringSubmatch(body)
	names := r.SubexpNames()
	if len(result) > 1 && len(names) > 1 {
		paramsMap := make(map[string]string)
		for i, name := range names {
			if i > 0 && i <= len(result) {
				paramsMap[name] = result[i]
			}
		}
		return paramsMap
	}
	return nil
}
