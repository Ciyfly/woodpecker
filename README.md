<!--
 * @Date: 2022-04-13 16:40:39
 * @LastEditors: recar
 * @LastEditTime: 2022-04-22 17:57:37
-->
## woodpecker 

啄木鸟  
可以使用xray nuclei yaml 以及go 代码的poc的poc扫描验证器

支持web后端模式和命令行模式
支持进度条数据(存入db中)  
支持yaml go poc导入db脚本  
支持存储xray的和gopoc的最后一次请求响应包入库当web模式

算是自己对cel-go的学习demo


## 使用  

### 编译
`go build -mod vendor cmd/woodpecker.go`  

### 命令行模式  
帮助信息  
![avatar](doc/imgs/help.jpg)  
开始扫描  
![avatar](doc/imgs/scan.jpg)  

### 接口模式

创建任务  
![avatar](doc/imgs/addtask.jpg)  
查看任务信息  
![avatar](doc/imgs/gettask.jpg)  

## poc to db
`python script/goscript2db.py`  
`python script/nucleipoc2db.py`  
`python script/xraypoc2db.py`  

## TODO
- 优化性能
- 存储nuclei的请求响应包
- 目前暂时去掉了 nuclei的fuzzing 和workflow的目录


参考  
https://github.com/WAY29/pocV   
https://github.com/jweny/pocassist  
https://github.com/chaitin/xray  
https://github.com/boy-hack/w14scan  
























































































