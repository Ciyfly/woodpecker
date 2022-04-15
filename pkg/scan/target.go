/*
 * @Date: 2022-04-15 17:27:27
 * @LastEditors: recar
 * @LastEditTime: 2022-04-15 17:57:00
 */
package scan

import (
	"net"
	"net/url"
	"sync"
	"time"
	"woodpecker/pkg/log"

	"github.com/panjf2000/ants/v2"
)

func Verify(target string) bool {
	log.Infof("Verifying target %s", target)
	parseUrl, err := url.Parse(target)
	if err != nil {
		log.Errorf("target parse error: %s", err.Error())
	}
	addrr := parseUrl.Host
	conn, err := net.DialTimeout("tcp", addrr, 2*time.Second)
	if err != nil {
		return false
	} else {
		conn.Close()
		return true
	}
}

func VerifyTargets(targets []string) []string {
	aliveTarget := []string{}
	var targetWg sync.WaitGroup
	var lock sync.Mutex
	targetPool, _ := ants.NewPoolWithFunc(50, func(data interface{}) {
		target := data.(string)
		verifyStatus := Verify(target)
		lock.Lock()
		if verifyStatus == true {
			aliveTarget = append(aliveTarget, target)
		}
		lock.Unlock()
		targetWg.Done()
	})
	defer targetPool.Release()
	for _, target := range targets {
		targetWg.Add(1)
		_ = targetPool.Invoke(target)
	}
	targetWg.Wait()
	return aliveTarget
}
