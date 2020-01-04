/*
@Created by : 2019/12/14 16:53
@Author : 何来亮
@Descripition :
*/
package pub

const ( //应用类型1：商户积分；2：商户点卡；3：零售票券
	AppTypeMerJF = 1
	AppTypeMerDK = 2
	AppTypePQ    = 3
)

const (
	MerStatusInit   = 1 //1：待审核；2：已开通；3；禁用；4：退网
	MerStatusOK     = 2
	MerStatusForbid = 3
	MerStatusDel    = 4
)

const (
	MerOprTypeForbid = 1 // 商户操作-禁用
	MerOprTypeDel    = 2 // 商户操作-删除
	MerOprTypeCheck  = 3 // 商户操作-审核
	MerOprTypeOpen   = 4 // 商户操作-开启
)

const (
	TermStatusInit   = 1 //1：待审核；2：已开通；3；禁用；4：退网
	TermStatusOK     = 2
	TermStatusForbid = 3
	TermStatusDel    = 4
)
const MerType_ZY = 1 //1-自有商户,2-银联商户
const MerType_YL = 2

const MerPayType_ZFB = 1 //1:支付宝; 2:银行
const MerPayType_YH = 2

const MerSettleType_ZD = 1 //1:自动2：手动
const MerSettleType_SD = 2

const LocalDateTimeFormat string = "2006-01-02 15:04:05"

const (
	PageNum  = 1  //默认的分页 页码
	PageSize = 10 // 每页的条数
)
