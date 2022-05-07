/*
 * @Date: 2022-04-26 10:53:25
 * @LastEditors: recar
 * @LastEditTime: 2022-04-27 16:21:34
 */
package scripts

import (
	"net/http"
	"woodpecker/pkg/log"
)

func VerifyTest2(sp *ScriptParams) bool {
	sp.PocName = "test2"
	sp.PocLink = "test2.test2.com"
	sp.Description = "test2"
	resp, err := http.Get(sp.Target + "/recar")
	if err != nil {
		log.Errorf("error: %s", err.Error())
	}
	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}

}

func init() {
	RegisterScript("test2", VerifyTest2)
}
