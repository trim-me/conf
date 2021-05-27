package src

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

//调用结构体
type RedisType struct {
	Db   interface{}
	Conn *redis.Pool
}

var RedisConfMap RedisConnConfMap

//conn list
type RedisConnConfMap struct {
	Redis map[string]*RedisConnConf `toml:"redis"`
}

//redis
type RedisConnConf struct {
	Host        string
	Password    string
	Port        uint
	MaxIdleConn int
	MaxOpenConn int
}

//初始化
func (redisConf *RedisConnConf) RedisInit() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     redisConf.MaxIdleConn,
		MaxActive:   redisConf.MaxOpenConn,
		IdleTimeout: 240 * time.Second,
		Wait:        false, //超过最大连接数处理方式(true=等待,false=抛异常)
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(3000)*time.Millisecond), //超时时间
				redis.DialDatabase(0),
				//redis.DialPassword(""), //放到了下面进行密码验证
			)
			if err != nil {
				log.Fatal("redis数据库连接失败:", err)
			}
			if redisConf.Password != "" {
				if _, err := conn.Do("AUTH", redisConf.Password); err != nil {
					conn.Close()
					log.Fatal("redis数据库授权失败:", err)
				}
			}
			return conn, nil
		},
	}
}

//init toml
func (redisConf *RedisConnConf) InitRedisToml(file string) {
	if _, err := toml.DecodeFile(FileStat(file), &RedisConfMap); err != nil {
		log.Fatal("init mysql failed:", err)
	}
}

//get conn pool
func (redisConf *RedisConnConf) GetPool(tomlName string) *redis.Pool {
	if _, ok := RedisConfMap.Redis[tomlName]; !ok {
		log.Fatal("no have database tomlName pool:", tomlName)
	}
	return RedisConfMap.Redis[tomlName].RedisInit()
}
