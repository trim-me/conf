package src

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"item/public/config"
	"item/public/logs"
	"log"
	"time"
)

//调用关键词
var Redis RedisConn

//调用结构体
type RedisConn struct{}

//连接池
var RedisPool *redis.Pool

//redis
type RedisType struct {
	Host        string
	Password    string
	Port        uint
	MaxIdleConn int
	MaxOpenConn int
}

//初始化
func Redisinit(conf*RedisType) {
	var redisConf = config.Redis
	RedisPool = &redis.Pool{
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

//Set 插入k,v值
func (r RedisConn) Set(k, v interface{}, expire ...int) (err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if _, err = conn.Do("SET", k, v); err != nil {
		logs.Log.Println("插入失败", err)
		return
	}
	if len(expire) > 0 {
		if _, err = conn.Do("EXPIRE", k, expire[0]); err != nil {
			logs.Log.Println("设置超时时间失败", err)
			return
		}
	}
	return
}

//Get
func (r RedisConn) Get(k interface{}) (value string, err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if value, err = redis.String(conn.Do("GET", k)); err != nil {
		logs.Log.Println("查询失败", err)
		return
	}
	return
}

//Del
func (r RedisConn) Del(k interface{}) (err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if _, err = conn.Do("DEL", k); err != nil {
		fmt.Println("删除失败", err)
		return err
	}
	return
}

//HSet
func (r RedisConn) HSet(name, k, v interface{}) (err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if _, err = conn.Do("HSET", name, k, v); err != nil {
		return
	}
	return
}

//HGet
func (r RedisConn) HGet(name, k interface{}) (value interface{}, err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if value, err = conn.Do("HGet", name, k); err != nil {
		return
	}
	return
}

//HGetAll
func (r RedisConn) HGetAll(name interface{}) (value map[string]string, err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if value, err = redis.StringMap(conn.Do("hgetall", name)); err != nil {
		return
	}
	return
}

//HDel 删除hName表里的k下表数据
func (r RedisConn) HDel(name, k interface{}) (err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if _, err = redis.String(conn.Do("HDEL", name, k)); err != nil {
		return
	}
	return
}

//HKeys 获取hash表的所有下表
func (r RedisConn) HKeys(hName interface{}) (value []string, err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if value, err = redis.Strings(conn.Do("HKEYS", hName)); err != nil {
		return
	}
	return
}

func (r RedisConn) RPush(k string, v interface{}) (err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if _, err = conn.Do("rpush", k, v); err != nil {
		return
	}
	return
}
func (r RedisConn) LPush(k string, v interface{}) (err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if _, err = conn.Do("lpush", k, v); err != nil {
		return
	}
	return
}
func (r RedisConn) LPop(k string) (reply string, err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if reply, err = redis.String(conn.Do("lpop", k)); err != nil {
		return
	}
	return
}
func (r RedisConn) RPop(k string) (reply string, err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if reply, err = redis.String(conn.Do("rpop", k)); err != nil {
		return
	}
	return
}

//遍历
func (r RedisConn) LRang(k string) (reply []int, err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	reply, err = redis.Ints(conn.Do("lrange", k, 0, -1))
	return
}

//删除
func (r RedisConn) LRem(k string, v interface{}) (err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	_, err = conn.Do("lrem", k, 2, v)
	return
}

func (r RedisConn) IncrBy(k string, v interface{}) (reply interface{}, err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	if reply, err = conn.Do("incrBy", k, v); err != nil {
		return
	}
	return
}
