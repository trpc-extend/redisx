package redisx

import (
	"fmt"
	"sync"
	"time"

	pcgmonitor "git.code.oa.com/pcgmonitor/trpc_report_api_go"
	"git.code.oa.com/pcgmonitor/trpc_report_api_go/pb/nmnt"
	"git.code.oa.com/trpc-go/trpc-go/errs"
)

var (
	once     sync.Once
	instance *pcgmonitor.Instance
)

// getInstance 获取007上报实例
func getInstance() *pcgmonitor.Instance {
	once.Do(func() {
		instance = pcgmonitor.NewInstance()
		_ = instance.CommSetUp(&pcgmonitor.CommSvrSetupInfo{
			CommSvrInfo: pcgmonitor.CommSvrInfo{
				CommName:  "weishi_component_redis",
				IP:        "",
				Container: "",
			},
		})
	})

	return instance
}

func commReport(dimensions []string, statValues []*nmnt.StatValue) {
	_ = getInstance().CommReport(dimensions, statValues)
}

func commReportWithSuffix(commItemName string, dimensions []string, statValues []*nmnt.StatValue) {
	_ = getInstance().CommReportWithSuffix(commItemName, dimensions, statValues)
}

func cmdReport(target, name, cmd string, ti time.Time, err error) {
	var dimensions = []string{
		ParseTarget(target),               // target,redis+polaris://:passwd@polaris_name,固定格式
		name,                              // serviceName,主调服务名
		RedisComponentVersion,             // 组件版本号
		cmd,                               // 命令名
		fmt.Sprintf("%d", errs.Code(err)), // 错误码
		errs.Msg(err),                     // 错误信息
	}

	succNum := err2num(err, ReturnSucc)
	timeCost := time.Since(ti).Milliseconds()
	faultNum := err2num(err, ReturnFault)
	timeoutNum := err2num(err, ReturnTimeout)
	cacheMissNum := err2num(err, ReturnCachemiss)
	connFaultNum := err2num(err, ReturnConnfailed)

	var statValues []*nmnt.StatValue
	statValues = append(statValues, &nmnt.StatValue{Value: 1, Policy: nmnt.Policy_SUM})                     //请求数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(succNum), Policy: nmnt.Policy_SUM})      //成功数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(timeCost), Policy: nmnt.Policy_AVG})     //耗时
	statValues = append(statValues, &nmnt.StatValue{Value: float64(faultNum), Policy: nmnt.Policy_SUM})     //异常数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(timeoutNum), Policy: nmnt.Policy_SUM})   //超时数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(cacheMissNum), Policy: nmnt.Policy_SUM}) //cachemiss数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(connFaultNum), Policy: nmnt.Policy_SUM}) //连接异常数

	commReportWithSuffix("cmd", dimensions, statValues)
}

// pipelineReport pipeline上报
func pipelineReport(target, name string, ti time.Time, err error, subcmd string, suberr error, statVal int) {
	var dimensions = []string{
		ParseTarget(target),                  // target,redis+polaris://:passwd@polaris_name,固定格式
		name,                                 // serviceName,主调服务名
		RedisComponentVersion,                // 组件版本号
		fmt.Sprintf("%d", errs.Code(err)),    // pipeline错误码
		errs.Msg(err),                        // pipeline错误信息
		subcmd,                               // 子命令名
		fmt.Sprintf("%d", errs.Code(suberr)), // 子命令错误码
		errs.Msg(suberr),                     // 子命令错误信息
	}

	succNum := err2num(err, ReturnSucc)
	timeCost := time.Since(ti).Milliseconds() //多次上报
	faultNum := err2num(err, ReturnFault)
	timeoutNum := err2num(err, ReturnTimeout)
	cacheMissNum := err2num(err, ReturnCachemiss)
	connFaultNum := err2num(err, ReturnConnfailed)

	if statVal == PipeLineCmdStatValZero {
		// pipeline本身上报是与子命令一起上报，只统计一次
		succNum = 0
		faultNum = 0
		timeoutNum = 0
		cacheMissNum = 0
		connFaultNum = 0
	}

	subSuccNum := err2num(suberr, ReturnSucc)
	subTimeCost := time.Since(ti).Milliseconds()
	subFaultNum := err2num(suberr, ReturnFault)
	subTimeoutNum := err2num(suberr, ReturnTimeout)
	subCacheMissNum := err2num(suberr, ReturnCachemiss)

	var statValues []*nmnt.StatValue
	// pipeLine 命令统计
	statValues = append(statValues, &nmnt.StatValue{Value: float64(statVal), Policy: nmnt.Policy_SUM})      //请求数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(succNum), Policy: nmnt.Policy_SUM})      //成功数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(timeCost), Policy: nmnt.Policy_AVG})     //耗时
	statValues = append(statValues, &nmnt.StatValue{Value: float64(faultNum), Policy: nmnt.Policy_SUM})     //异常数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(timeoutNum), Policy: nmnt.Policy_SUM})   //超时数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(cacheMissNum), Policy: nmnt.Policy_SUM}) //cachemiss数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(connFaultNum), Policy: nmnt.Policy_SUM}) //连接异常数

	// pipeline 子命令统计
	statValues = append(statValues, &nmnt.StatValue{Value: 1, Policy: nmnt.Policy_SUM})                      //子命令请求数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(subSuccNum), Policy: nmnt.Policy_SUM})    //子命令成功数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(subTimeCost), Policy: nmnt.Policy_AVG})   //子命令耗时
	statValues = append(statValues, &nmnt.StatValue{Value: float64(subFaultNum), Policy: nmnt.Policy_SUM})   //子命令异常数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(subTimeoutNum), Policy: nmnt.Policy_SUM}) //子命令超时数
	statValues = append(statValues, &nmnt.StatValue{Value: float64(subCacheMissNum), Policy: nmnt.Policy_SUM})

	commReportWithSuffix("pipeline", dimensions, statValues)
}
