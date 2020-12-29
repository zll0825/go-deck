package transport

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/forward/transport/lishou"
)

func InitTransportRouter(router *gin.RouterGroup) {
	app := router.Group("app")
	{
		app.POST("createAppBaseInfo", lishou.CreateAppBaseInfo)     // 添加机构基本信息
		app.POST("createAppFee", lishou.CreateAppFee)               // 配置成本费率
		app.POST("createAppDefaultFee", lishou.CreateAppDefaultFee) // 配置默认成本费率
		app.POST("createAppAttachInfo", lishou.CreateAppAttachInfo) // 添加机构附件信息
		app.POST("submitApp", lishou.SubmitApp)                     // 提交审核
		app.POST("appManageList", lishou.AppManageList)             // 机构管理列表（编辑表）
		app.POST("appManageDetail", lishou.AppManageDetail)         // 机构详情（编辑表）
		app.POST("removeApp", lishou.RemoveApp)                     // 注销机构
		app.POST("freezeApp", lishou.FreezeApp)                     // 冻结机构
		app.POST("appInfoList", lishou.AppInfoList)                 // 机构信息列表（主表）
		app.POST("appInfoDetail", lishou.AppInfoDetail)             // 机构信息详情（主表）
	}

	mcht := router.Group("mcht")
	{
		mcht.POST("mchtManageList", lishou.MchtManageList)       // 商户管理列表（编辑表）
		mcht.POST("mchtManageDetail", lishou.MchtManageDetail)   // 商户详情（编辑表）
		mcht.POST("removeMcht", lishou.RemoveMcht)               // 注销商户
		mcht.POST("freezeMcht", lishou.FreezeMcht)               // 冻结商户
		mcht.POST("mchtInfoList", lishou.MchtInfoList)           // 商户信息列表（主表）
		mcht.POST("mchtInfoDetail", lishou.MchtInfoDetail)       // 商户信息详情（主表）
		mcht.POST("mchtApplyList", lishou.MchtApplyList)         // 商户申请列表
		mcht.POST("mchtApplyDetail", lishou.MchtApplyDetail)     // 商户申请详情
		mcht.POST("mchtApplyWXConfig", lishou.MchtApplyWXConfig) // 商户申请管理-微信配置
	}

	term := router.Group("term")
	{
		term.POST("termInfoList", lishou.TermInfoList)     // 终端信息列表
		term.POST("termInfoDetail", lishou.TermInfoDetail) // 终端信息详情
	}

	trade := router.Group("trade")
	{
		trade.POST("todayTradeList", lishou.TodayTradeList)         // 当日流水列表
		trade.POST("todayTradeDetail", lishou.TodayTradeDetail)     // 当日流水详情
		trade.POST("historyTradeList", lishou.HistoryTradeList)     // 历史流水列表
		trade.POST("historyTradeDetail", lishou.HistoryTradeDetail) // 历史流水详情
		trade.POST("liveTradeList", lishou.LiveTradeList)           // 实时流水列表
		trade.POST("liveTradeDetail", lishou.LiveTradeDetail)       // 实时流水详情
	}

	settle := router.Group("settle")
	{
		settle.POST("d0RepayList", lishou.D0RepayList)       // D0补付列表
		settle.POST("d0RepayDetail", lishou.D0RepayDetail)   // D0补付详情
		settle.POST("d0RepayImport", lishou.D0RepayImport)   // D0补付导入
		settle.POST("d0RepayExport", lishou.D0RepayExport)   // D0补付导出
		settle.POST("reportDownload", lishou.ReportDownload) // 清算报表下载
	}

	risk := router.Group("risk")
	{
		risk.POST("createCashLimitRule", lishou.CreateCashLimitRule)   // 新增商户限额规则
		risk.POST("updateCashLimitRule", lishou.UpdateCashLimitRule)   // 修改商户限额规则
		risk.POST("cashLimitRuleDetail", lishou.CashLimitRuleDetail)   // 商户限额规则详情
		risk.POST("validCashLimitRule", lishou.ValidCashLimitRule)     // 生效商户限额规则
		risk.POST("invalidCashLimitRule", lishou.InvalidCashLimitRule) // 作废商户限额规则
		risk.POST("createInputRule", lishou.CreateInputRule)           // 新增商户进件规则
		risk.POST("updateInputRule", lishou.UpdateInputRule)           // 修改商户进件规则
		risk.POST("inputRuleDetail", lishou.InputRuleDetail)           // 商户进件规则详情
		risk.POST("validInputRule", lishou.ValidInputRule)             // 生效商户进件规则
		risk.POST("invalidInputRule", lishou.InvalidInputRule)         // 作废商户进件规则
		risk.POST("createInputWhite", lishou.CreateInputWhite)         // 新增商户进件白名单
		risk.POST("updateInputWhite", lishou.UpdateInputWhite)         // 修改商户进件白名单
		risk.POST("inputWhiteDetail", lishou.InputWhiteDetail)         // 商户进件白名单详情
		risk.POST("validInputWhite", lishou.ValidInputWhite)           // 生效商户进件白名单
		risk.POST("invalidInputWhite", lishou.InvalidInputWhite)       // 作废商户进件白名单
	}

	qr := router.Group("qr")
	{
		qr.POST("qrList", lishou.QrList)     // 聚合码列表
		qr.POST("qrConfig", lishou.QrConfig) // 聚合码配置
	}
}
