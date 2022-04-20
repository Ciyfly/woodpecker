/*
 * @Date: 2022-04-18 18:22:50
 * @LastEditors: recar
 * @LastEditTime: 2022-04-20 11:41:05
 */
package parse

// 扫描结果
const (
	ResultFail = iota
	ResultSuccess
	ResultError
)

// 任务状态
const (
	ScanCreate = iota
	ScanRun
	ScanEnd
	ScanError
)

// mode 扫描模式

const (
	ModeServer = "server"
	ModeScan   = "scan"
)

// poc开启状态
const (
	PocNotEnable = iota
	PocEnable
)
