package src

//
////Set 插入k,v值
//func (r RedisType) Set(k, v interface{}, expire ...int) (err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if _, err = conn.Do("SET", k, v); err != nil {
//		return
//	}
//	if len(expire) > 0 {
//		if _, err = conn.Do("EXPIRE", k, expire[0]); err != nil {
//			return
//		}
//	}
//	return
//}
//
////Get
//func (r RedisType) Get(k interface{}) (value string, err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if value, err = redis.String(conn.Do("GET", k)); err != nil {
//		return
//	}
//	return
//}
//
////Del
//func (r RedisType) Del(k interface{}) (err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if _, err = conn.Do("DEL", k); err != nil {
//		fmt.Println("删除失败", err)
//		return err
//	}
//	return
//}
//
////HSet
//func (r RedisType) HSet(name, k, v interface{}) (err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if _, err = conn.Do("HSET", name, k, v); err != nil {
//		return
//	}
//	return
//}
//
////HGet
//func (r RedisType) HGet(name, k interface{}) (value interface{}, err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if value, err = conn.Do("HGet", name, k); err != nil {
//		return
//	}
//	return
//}
//
////HGetAll
//func (r RedisType) HGetAll(name interface{}) (value map[string]string, err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if value, err = redis.StringMap(conn.Do("hgetall", name)); err != nil {
//		return
//	}
//	return
//}
//
////HDel 删除hName表里的k下表数据
//func (r RedisType) HDel(name, k interface{}) (err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if _, err = redis.String(conn.Do("HDEL", name, k)); err != nil {
//		return
//	}
//	return
//}
//
////HKeys 获取hash表的所有下表
//func (r RedisType) HKeys(hName interface{}) (value []string, err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if value, err = redis.Strings(conn.Do("HKEYS", hName)); err != nil {
//		return
//	}
//	return
//}
//
//func (r RedisType) RPush(k string, v interface{}) (err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if _, err = conn.Do("rpush", k, v); err != nil {
//		return
//	}
//	return
//}
//func (r RedisType) LPush(k string, v interface{}) (err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if _, err = conn.Do("lpush", k, v); err != nil {
//		return
//	}
//	return
//}
//func (r RedisType) LPop(k string) (reply string, err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if reply, err = redis.String(conn.Do("lpop", k)); err != nil {
//		return
//	}
//	return
//}
//func (r RedisType) RPop(k string) (reply string, err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if reply, err = redis.String(conn.Do("rpop", k)); err != nil {
//		return
//	}
//	return
//}
//
////遍历
//func (r RedisType) LRang(k string) (reply []int, err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	reply, err = redis.Ints(conn.Do("lrange", k, 0, -1))
//	return
//}
//
////删除
//func (r RedisType) LRem(k string, v interface{}) (err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	_, err = conn.Do("lrem", k, 2, v)
//	return
//}
//
//func (r RedisType) IncrBy(k string, v interface{}) (reply interface{}, err error) {
//	conn := r.Conn.Get()
//	defer conn.Close()
//	if _, err := conn.Do("SELECT", r.Db); err != nil {
//		return
//	}
//	if reply, err = conn.Do("incrBy", k, v); err != nil {
//		return
//	}
//	return
//}
