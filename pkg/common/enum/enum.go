/*
 * @Date: 2022-04-18 18:22:50
 * @LastEditors: recar
 * @LastEditTime: 2022-04-27 11:02:53
 */
package enum

// 扫描结果
const (
	ResultFail = iota
	ResultSuccess
	ResultError
)

// 验证类型
const (
	XrayVerify = iota
	NuclieVerify
	GoScriptVerify
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
