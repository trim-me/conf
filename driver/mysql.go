package driver

import (
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MysqlConfMapList MysqlConnConfMap

type MysqlConnConfMap struct {
	Mysql map[string]*MysqlConnConf `toml:"mysql"`
}

type MysqlConnConf struct {
	DriverName      string `toml:"driver_name"`
	Host            string `toml:"host"`
	Database        string `toml:"database"`
	User            string `toml:"user"`
	Password        string `toml:"password"`
	Port            uint   `toml:"port"`
	MaxIdleConn     int    `toml:"max_idle_conn"`      //用于设置闲置的连接数。
	MaxOpenConn     int    `toml:"max_open_conn"`      //用于设置最大打开的连接数，默认值为0表示不限制
	ConnMaxIdleTime int    `toml:"conn_max_idle_time"` //链接空闲时长 超出杀死
	ConnMaxLifetime int    `toml:"conn_max_life_time"` //链接复用时长
	SqlLog          int    `toml:"sql_log"`            //日志级别 1 静默 2 错误信息 3 警告信息 4 所有信息
}

//open conn
func (conf *MysqlConnConf) mysqlInit() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=False&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database)
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("mysql connect failed:", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("mysql connect failed:", err)
	}
	sqlDB.SetMaxIdleConns(conf.MaxIdleConn) //max free(idle) conn
	sqlDB.SetMaxOpenConns(conf.MaxOpenConn) //max conn
	sqlDB.SetConnMaxIdleTime(time.Minute)   //max idle conn time
	sqlDB.SetConnMaxLifetime(time.Hour)
	db.Logger.LogMode(logger.LogLevel(conf.SqlLog)) //sql run log whether output
	return db
}

func (conf *MysqlConnConf) InitMysqlToml(file string) {
	if _, err := toml.DecodeFile(FileStat(file), &MysqlConfMapList); err != nil {
		log.Fatal("init mysql failed:", err)
	}
}

//get database by name
func (conf *MysqlConnConf) GetPool(tomlName string) *gorm.DB {
	if _, ok := MysqlConfMapList.Mysql[tomlName]; !ok {
		log.Fatal("no have database tomlName pool:", tomlName)
	}
	return MysqlConfMapList.Mysql[tomlName].mysqlInit()
}
