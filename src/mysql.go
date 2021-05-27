package src

import (
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var	MysqlConfMapList MysqlConnConfMap

type MysqlConnConfMap struct {
	Mysql map[string]*MysqlConnConf `toml:"mysql"`
}

type MysqlConnConf struct {
	DriverName  string `toml:"driver_name"`
	Host        string `toml:"host"`
	Database    string `toml:"database"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
	Port        uint   `toml:"port"`
	MaxIdleConn int    `toml:"max_idle_conn"` //用于设置闲置的连接数。
	MaxOpenConn int    `toml:"max_open_conn"` //用于设置最大打开的连接数，默认值为0表示不限制
	SqlLog      bool   `toml:"sql_log"`       //是否打印sql
}

//初始化
func (conf *MysqlConnConf) mysqlInit() *gorm.DB {
	db, err := gorm.Open(
		conf.DriverName,
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
		log.Fatal("mysql connect failed:", err)
	}
	//连接池信息
	db.DB().SetMaxIdleConns(conf.MaxIdleConn) //设置最大空闲数
	db.DB().SetMaxOpenConns(conf.MaxOpenConn) //设置最大连接数
	db.SingularTable(true)
	db.LogMode(conf.SqlLog) //请求日志
	return db
}

func (conf *MysqlConnConf) InitMysqlToml(file string) {
	if _, err := toml.DecodeFile(FileStat(file), &MysqlConfMapList); err != nil {
		log.Fatal("init mysql failed:", err)
	}
}

//获取指定数据库连接池
func (conf *MysqlConnConf) GetPool(tomlName string) *gorm.DB {
	if _, ok := MysqlConfMapList.Mysql[tomlName]; !ok {
		log.Fatal("no have database tomlName pool:", tomlName)
	}
	return MysqlConfMapList.Mysql[tomlName].mysqlInit()
}
