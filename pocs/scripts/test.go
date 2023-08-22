/*
 * @Date: 2022-04-26 10:53:25
 * @LastEditors: recar
 * @LastEditTime: 2022-04-27 16:22:01
 */
package scripts

import (
	"net/http"
	"woodpecker/pkg/log"
)

func VerifyTest(sp *ScriptParams) bool {
	sp.PocName = "test"
	sp.PocLink = "test.test.com"
	sp.Description = "test"
	resp, err := http.Get(sp.Target)
	if err != nil {
		log.Errorf("error: %s", err.Error())
		return false
	}
	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}

}

func init() {
	RegisterScript("test", VerifyTest)
}
