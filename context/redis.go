package context

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
)

//Redis redia连接对象容器
type Redis struct {
	Pool redisConn
}

//redisConn 第三连接池
type redisConn struct {
	redis.Conn
}

func (r redisConn) Close() {
	r.Conn.Close()
}

//NewRedisPool 返回一个redis连接池
func NewRedisPool(o *Options, l *Log) *Redis {
	p := pools.NewResourcePool(func() (pools.Resource, error) {
		var c redis.Conn
		var err error
		if o.RedisAuth != "" {
			c, err = redis.Dial("tcp", o.RedisAddress, redis.DialPassword(o.RedisAuth))
		} else {
			c, err = redis.Dial("tcp", o.RedisAddress)
		}
		return redisConn{c}, err
	}, 1, 2, time.Minute)
	//defer p.Close()
	ctx := context.TODO()
	r, err := p.Get(ctx)
	if err != nil {
		l.Fatal("Redis连接池初始化错误", err)
	}
	defer p.Put(r)
	c := r.(redisConn)
	redis := new(Redis)
	redis.Pool = c
	return redis
}

//ScanStruct ScanStruct
func (r *Redis) ScanStruct(src []interface{}, dest interface{}) error {
	return redis.ScanStruct(src, dest)
}

//HMSet hash 批量设置
func (r *Redis) HMSet(key string, args interface{}) error {
	if key == "" {
		return errors.New("缺少key")
	}
	_, err := r.Pool.Do("HMSET", redis.Args{}.Add(key).AddFlat(args)...)
	return err
}

//HGetAll hash 批量获取
func (r *Redis) HGetAll(key string, container interface{}) error {
	data, err := redis.Values(r.Pool.Do("HGETALL", key))
	if err != nil {
		return err
	}
	err = redis.ScanStruct(data, container)
	return err
}
