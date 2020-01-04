package utils

const (
	RECODE_OK         = "0000"
	RECODE_DBERR      = "5001"
	RECODE_NODATA     = "5002"
	RECODE_DATAEXIST  = "5003"
	RECODE_DATAERR    = "5004"
	RECODE_SESSIONERR = "5005"
	RECODE_LOGINERR   = "5006"
	RECODE_PARAMERR   = "5007"
	RECODE_USERERR    = "5008"
	RECODE_ROLEERR    = "5009"
	RECODE_PWDERR     = "5010"
	RECODE_SMSERR     = "5011"
	RECODE_REQERR     = "5012"
	RECODE_IPERR      = "5013"
	RECODE_THIRDERR   = "5014"
	RECODE_IOERR      = "5015"
	RECODE_SERVERERR  = "5016"
	RECODE_UNKNOWERR  = "5017"

	RECODE_EXIST_MER_FAIL  = "5018"
	RECODE_EXIST_TERM_FAIL = "5019"
)

var recodeText = map[string]string{
	RECODE_OK:              "成功",
	RECODE_DBERR:           "数据库查询错误",
	RECODE_NODATA:          "无数据",
	RECODE_DATAEXIST:       "数据已存在",
	RECODE_DATAERR:         "数据错误",
	RECODE_SESSIONERR:      "用户未登录",
	RECODE_LOGINERR:        "用户登录失败",
	RECODE_PARAMERR:        "参数错误",
	RECODE_USERERR:         "用户不存在或未激活",
	RECODE_ROLEERR:         "用户身份错误",
	RECODE_PWDERR:          "密码错误",
	RECODE_REQERR:          "非法请求或请求次数受限",
	RECODE_IPERR:           "IP受限",
	RECODE_THIRDERR:        "第三方系统错误",
	RECODE_IOERR:           "文件读写错误",
	RECODE_SERVERERR:       "内部错误",
	RECODE_UNKNOWERR:       "未知错误",
	RECODE_SMSERR:          "短信失败",
	RECODE_EXIST_MER_FAIL:  "无此商户信息",
	RECODE_EXIST_TERM_FAIL: "无此终端信息",
}

func GetErrMsg(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
