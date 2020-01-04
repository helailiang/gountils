/*
 @ Author : 高侃学
 @ Time : 2019/6/27 15:43
 @ Describe :
*/
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fiveCity/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
 @ Author : 高侃学
 @ Describe : 获取随机数
 @ InParam : 位数（6位或8位）
*/
func RandAmounts(k int) string {
	var code string
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if k == 6 {
		code = fmt.Sprintf("%06v", rnd.Int31n(1000000))
	} else {
		code = fmt.Sprintf("%08v", rnd.Int31n(100000000))
	}
	return code
}

/*
 @ Author : 高侃学
 @ Describe : 获取当前日期和时间
 @ InParam : d-日期分隔符；t-时间分隔符
*/
func GetDate(d, t string) (string, string) {
	now := time.Now()
	curDate := fmt.Sprintf("%4d%s%02d%s%02d", now.Year(), d, now.Month(), d, now.Day())
	curTime := fmt.Sprintf("%02d%s%02d%s%02d", now.Hour(), t, now.Minute(), t, now.Second())
	return curDate, curTime
}

/**
 * @Author 何来亮
 * @Date 17:03 2019/12/13
 * @Description md5加密
 * @Param
 * @return
 **/
func MD5(str string) string {
	hs := md5.New()
	if _, err := hs.Write([]byte(str)); err != nil {
		logger.Log.Panic(err)
	}
	return strings.ToUpper(hex.EncodeToString(hs.Sum(nil)))
}

/**
 * @Author 何来亮
 * @Date 16:26 2019/7/12
 * @Description 获取本地ip
 * @Param
 * @return
 **/
func FetchIP() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for index := range addrs {

		// 检查ip地址判断是否回环地址
		if IPNet, ok := addrs[index].(*net.IPNet); ok && !IPNet.IP.IsLoopback() {
			if IPNet.IP.To4() != nil {
				return IPNet.IP, nil
			}

		}
	}

	return nil, errors.New("failed to found IP address")
}

/**
 * @Author 何来亮
 * @Date 17:14 2019/7/12
 * @Description  生成唯一的序列号 YYYYmmddHHMMSS+时间微妙
 * @Param
 * @return string
 **/
func GetSeqNumber() string {
	//strconv.FormatInt int64到string
	return time.Now().Format("20060102150405") + strconv.FormatInt(time.Now().UnixNano()/1000, 10)

}

/**
 * @Author 何来亮
 * @Date 9:54 2019/7/15
 * @Description 微信支付计算签名的函数,所有参数按照字段名的ascii码从小到大排序,进行md5加密
 * @Param mReq：加密原始数据   key:商户密钥
 * @return 加密的md5串
 **/
func WxPayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		logger.Log.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	fmt.Println("加密前-----", signStrings)
	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings)) //
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))

	fmt.Println("加密后-----", upperSign)
	return upperSign
}

/**
 * @Author 何来亮
 * @Date 14:12 2019/7/16
 * @Description  获取当前时间字符串 格式为YYYYmmDDHHMMSS
 * @Param
 * @return string
 **/
func GetDataTimeStr() string {
	return time.Now().Format("20060102150405")

}

/**
 * @Author 何来亮
 * @Date 12:09 2019/12/14
 * @Description  将时间字符串转换成时间
 * @Param
 * @return
 **/
func GetStrToDataTime(tStr string) (time.Time, error) {
	t, err := time.Parse("20060102150405", tStr)
	if err != nil {
		logger.Log.Error(err)
	}
	return t, err
}

/**
 * @Author 何来亮
 * @Date 12:09 2019/12/14
 * @Description  将时间字符串转换成日期
 * @Param
 * @return
 **/
func GetStrToData(tStr string) time.Time {
	t, err := time.Parse("20060102", tStr)
	if err != nil {
		logger.Log.Panic(err)
	}
	return t
}

/**
 * @Author 何来亮
 * @Date 11:32 2019/8/14
 * @Description  获取当前时间半小时之前的时间字符串
 * @Param
 * @return
 **/
func GetNowBeforeOneHour() string {
	return time.Now().Add(-time.Minute * 30).Format("20060102150405")
}

/**
 * @Author 何来亮
 * @Date 15:13 2019/7/25
 * @Description 将日期格式化 20060102150405转为2006-01-02 15:04
 * @Param
 * @return
 **/
func StrDataToFormat(strData string) string {
	if len(strData) == 0 {
		strData = time.Now().Format("20060102150405")
	}
	t, _ := time.Parse("20060102150405", strData)
	return t.Format("2006-01-02 15:04")
}

/**
 * @Author 何来亮
 * @Date 14:11 2019/7/23
 * @Description 将元转分
 * @Param
 * @return
 **/
func YuanToFen(yuan string) string {
	i, _ := strconv.ParseFloat(yuan, 32)
	return fmt.Sprintf("%.f", i*100)
}

/**
 * @Author 何来亮
 * @Date 14:02 2019/7/25
 * @Description  将分转元
 * @Param
 * @return
 **/
func FenToYuan(fen string) string {
	i, _ := strconv.ParseFloat(fen, 32)
	return fmt.Sprintf("%.2f", i/100)
}

/**
 * @Author 何来亮
 * @Date 18:37 2019/8/5
 * @Description
 * @Param  判断手机号所属运行商
 * @return 1-中国移动 2-中国联通  3-中国电信 0-查询出错
 **/
func JudgeMobileOperator(mobile string) (MType int, err error) {
	// 为了配合测试话费包，先将135 转为联通
	CN := []int{134, 136, 137, 138, 139, 150, 151, 152, 157, 158, 159, 182,
		183, 184, 187, 188, 147, 178}
	CNUnion := []int{130, 131, 132, 155, 156, 185, 186, 145, 176, 179, 135}
	CNDianxin := []int{133, 153, 180, 181, 189, 177}
	nums, _ := strconv.Atoi(mobile[:3])
	err = nil
	if len([]rune(mobile)) == 11 {
		if IsExiestInArray(nums, CN) {
			logger.Log.Info("中国移动")
			MType = 1
		} else if IsExiestInArray(nums, CNUnion) {
			logger.Log.Info("中国联通")
			MType = 2
		} else if IsExiestInArray(nums, CNDianxin) {
			logger.Log.Info("中国电信")
			MType = 3
		} else {
			logger.Log.Info("查不到归属运营商")
			MType = 0
			err = errors.New("查不到归属运营商")
		}
	} else {
		logger.Log.Info("输入错误，请输入11位数字")
		MType = 0
		err = errors.New("输入的手机号有误，请输入有效的手机号")
	}
	return
}

/**
 * @Author 何来亮
 * @Date 18:38 2019/8/5
 * @Description
 * @Param 判断item是否在数组中，只能是[]int类型
 * @return
 **/
func IsExiestInArray(value int, arr []int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false

}

/**
 * @Author 何来亮
 * @Date 15:11 2019/8/7
 * @Description  判断slice是否存在某个item,可以是任意类型的切片
 * @Param
 * @return
 **/

func IsExistItem(value interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() { //TypeOf是返回[]int []string Kind返回该接口的具体分类 slice
	case reflect.Slice: // 若是切片
		s := reflect.ValueOf(array) // 返回反射Value，
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}

	}
	return false
}

type PublicResErr struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

/**
 * @Author 何来亮
 * @Date 17:27 2019/8/6
 * @Description
 * @Param 捕获异常 类似try catch,在文件头声明,并且需要返回报文的时候
 * @return
 **/
func CoverErrorMessage(c *gin.Context) {
	if message := recover(); message != nil {
		var err error
		switch x := message.(type) {
		case PublicResErr:
			res, _ := json.Marshal(message)
			logger.Log.Error("返回的响应报文为：", string(res))
			c.JSON(http.StatusOK, message)
			return
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("未定义异常")
		}
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		logger.Log.Info("Recovered panic error : %s  ==>%s\n ", err, string(buf[:n]))
		c.JSON(http.StatusOK, gin.H{
			"code": RECODE_SERVERERR,
			"msg":  fmt.Sprintf("%s:%s", GetErrMsg(RECODE_SERVERERR), err),
		})
		return
	}
}

/**
 * @Author 何来亮
 * @Date 13:51 2019/8/9
 * @Description 仅仅捕获异常
 * @Param
 * @return
 **/
func CoverErrorMessageOnly() {
	if message := recover(); message != nil {
		var err error
		switch x := message.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("未定义异常")
		}
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		logger.Log.Error("Recovered panic error : %s  ==>%s\n ", err, string(buf[:n]))
	}
}

/**
 * @Author 何来亮
 * @Date 9:54 2019/9/9
 * @Description  拼接指定次数的重复字符串，
 * @Param
 * @return
 **/
func RepeatStr(repStr string, num int) (resStr string) {
	for i := 0; i < num; i++ {
		resStr = strings.Join([]string{resStr, repStr}, "")
	}
	return
}

/**
 * @Author 何来亮
 * @Date 14:28 2019/12/13
 * @Description   获取UUID
 * @Param
 * @return
 **/
func GetUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
