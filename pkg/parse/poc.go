/*
 * @Date: 2022-04-13 17:17:01
 * @LastEditors: recar
 * @LastEditTime: 2022-04-14 11:26:53
 */
package parse

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
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
	Version     string   `yaml:"version"`
	Tags        string   `yaml:"tags"`
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
		log.Errorf("Error loading YAML file: %s", err)
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
			log.Errorf("Failed to read poc file %s %s", pocPath, err)
		}
		err = yaml.Unmarshal(yamlFile, poc)
		if err != nil {
			log.Errorf("Failed to unmarshal poc file %s %s", pocPath, err)
		}
		pocList = append(pocList, poc)
	}
	return pocList
}
