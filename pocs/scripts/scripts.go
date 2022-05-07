/*
 * @Date: 2022-04-26 10:55:09
 * @LastEditors: recar
 * @LastEditTime: 2022-04-27 17:22:11
 */
package scripts

import (
	"net/http"
	"woodpecker/pkg/db"
)

type ScriptParams struct {
	Target      string
	TaskId      int
	PocName     string
	PocLink     string
	Description string
	Mode        string
	Req         *http.Request
	Rsp         *http.Response
}

type ScriptFunc func(sp *ScriptParams) bool

var ScriptFuncMap = map[string]ScriptFunc{}

func RegisterScript(pocName string, ScriptFunction ScriptFunc) {
	ScriptFuncMap[pocName] = ScriptFunction
}

//go script

func LoadScriptPoc(pocNames []string) []ScriptFunc {
	goScriptList := []ScriptFunc{}
	if len(pocNames) == 0 {
		for name := range ScriptFuncMap {
			goScriptList = append(goScriptList, ScriptFuncMap[name])
		}
		return goScriptList
	}
	for _, name := range pocNames {
		if scriptFunc, ok := ScriptFuncMap[name]; ok {
			goScriptList = append(goScriptList, scriptFunc)
		}
	}
	return goScriptList
}

func Db2ScriptPoc(dbPocs []db.Poc) []ScriptFunc {
	goScriptList := []ScriptFunc{}
	for _, poc := range dbPocs {
		if scriptFunc, ok := ScriptFuncMap[poc.PocName]; ok {
			goScriptList = append(goScriptList, scriptFunc)
		}
	}
	return goScriptList
}
