package src

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var mysqlConf *MysqlType

type MysqlType struct {
	Host        string
	Database    string
	User        string
	Password    string
	Port        uint
	MaxIdleConn int  //用于设置闲置的连接数。
	MaxOpenConn int  //用于设置最大打开的连接数，默认值为0表示不限制
	SqlLog      bool //是否打印sql
}

//初始化配置文件
func init() {
	var (
		path            string
		str, _          = os.Getwd()
		defaultConfPath = str + "/conf/api/mysql.toml"
	)
	_, err := os.Stat(defaultConfPath)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}
	//设置模式
	flag.StringVar(&path, "conf-path", defaultConfPath, "加载Mysql配置")
	flag.Parse()
	if _, err := toml.DecodeFile(path, &mysqlConf); err != nil {
		log.Fatal("初始化配置失败:", err)
	}
	mysqlConf.mysqlInit()
}

//初始化
func (conf *MysqlType) mysqlInit() *gorm.DB {
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=False&loc=Local",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Database,
		),
	)
	if err != nil {
		log.Fatal("mysql数据库连接失败:", err)
	}
	//连接池信息
	db.DB().SetMaxIdleConns(conf.MaxIdleConn) //设置最大空闲数
	db.DB().SetMaxOpenConns(conf.MaxOpenConn) //设置最大连接数
	db.SingularTable(true)
	db.LogMode(conf.SqlLog) //请求日志
	return db
}
