package utils

import (
	"fiveCity/logger"
	//使用了beego框架的配置文件读取模块
	"github.com/astaxie/beego/config"
)

var (
	G_server_name  string //项目名称
	G_server_addr  string //服务器ip地址
	G_server_port  string //服务器端口
	G_read_timeout int
	G_write_timeout int
	G_redis_addr   string //redis ip地址
	G_redis_port   string //redis port端口
	G_redis_dbnum  string //redis db 编号
	G_mysql_addr   string //mysql ip 地址
	G_mysql_port   string //mysql 端口
	G_mysql_dbname string //mysql db name
	G_fastdfs_port string //fastdfs 端口
	G_fastdfs_addr string //fastdfs ip
	G_orgCode 	   string //组织编号,用于生成商户号和终端号
	G_merAppPwd    string //商户APP初始登陆密码
)

func InitConfig() {
	//从配置文件读取配置信息
	//如果项目迁移需要进行修改
	appconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		logger.Log.Error(err)
		return
	}
	G_server_name = appconf.String("appname")
	G_server_addr = appconf.String("httpaddr")
	G_server_port = appconf.String("httpport")
	G_read_timeout,_ = appconf.Int("ReadTimeout")
	G_write_timeout,_ = appconf.Int("WriteTimeout")
	G_redis_addr = appconf.String("redisaddr")
	G_redis_port = appconf.String("redisport")
	G_redis_dbnum = appconf.String("redisdbnum")
	G_mysql_addr = appconf.String("mysqladdr")
	G_mysql_port = appconf.String("mysqlport")
	G_mysql_dbname = appconf.String("mysqldbname")
	G_fastdfs_port = appconf.String("fastdfsport")
	G_fastdfs_addr = appconf.String("fastdfsaddr")
	G_orgCode = appconf.String("orgCode")
	G_merAppPwd = appconf.String("merAppPwd")
	return
}

func init() {
	InitConfig()
}
