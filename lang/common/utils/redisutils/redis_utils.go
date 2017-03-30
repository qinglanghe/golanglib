package redisutils

import (
	"github.com/garyburd/redigo/redis"
)

func GetUint64(pool *redis.Pool,key string) uint64 {
	rc := pool.Get()
	defer rc.Close()

	v , _:= redis.Uint64(rc.Do("GET",key))
	return v
}

func GetString(pool *redis.Pool,key string) string {
	rc := pool.Get()
	defer rc.Close()

	v , _:= redis.String(rc.Do("GET",key))
	return v
}

/**
* timeSecond = -1 表示永久
 */
func Set(pool *redis.Pool,key string,value interface{},timeSecond int64)  {
	rc := pool.Get()
	defer rc.Close()

	rc.Do("SET",key,value)
	if timeSecond != -1{
		rc.Do("EXPIRE",key,timeSecond)
	}
}

func SetExpire(pool *redis.Pool,key string,timeSecond int64)  {
	rc := pool.Get()
	defer rc.Close()

	rc.Do("EXPIRE",key,timeSecond)
}

func GetExpire(pool *redis.Pool,key string) int64 {
	rc := pool.Get()
	defer rc.Close()
	v , _:= redis.Int64(rc.Do("TTL",key))
	return v
}

func Del(pool *redis.Pool,key string)  {
	rc := pool.Get()
	defer rc.Close()
	rc.Do("DEL",key)
}

func Exists(pool *redis.Pool,key string) bool {
	rc := pool.Get()
	defer rc.Close()

	v , _ := redis.Bool(rc.Do("EXISTS",key))
	return v
}

func Incr(pool *redis.Pool,key string) uint64 {
	rc := pool.Get()
	defer rc.Close()
	v , _ := redis.Uint64(rc.Do("INCR",key))
	return v
}