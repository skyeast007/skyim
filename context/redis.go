package context

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
)

//Redis redia连接对象容器
type Redis struct {
	Redis pools.Resource
}

//RedisConn 热第三连接池
type RedisConn struct {
	redis.Conn
}

func (r RedisConn) Close() {
	r.Conn.Close()
}

//Resource 返回一个redis连接池
func NewRedisPool(o *Options, l *Log) *Redis {
	p := pools.NewResourcePool(func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", o.RedisAddress)
		return RedisConn{c}, err
	}, 1, 2, time.Minute)
	defer p.Close()
	ctx := context.TODO()
	r, err := p.Get(ctx)
	if err != nil {
		l.Fatal("redis连接池初始化错误", err)
	}
	defer p.Put(r)
	c := r.(RedisConn)
	if o.RedisAuth != "" {
		_, err = c.Do("auth", o.RedisAuth)
		if err != nil {
			l.Fatal("redis授权失败", err)
		}
	}
	redis := new(Redis)
	redis.Redis = c
	return redis
}
