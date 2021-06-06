package redis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	pool      *redis.Pool
	redisHost = "127.0.0.1:6379"
)

//newRedisPool：创建redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		//连接池中最多有多少条连接可以使用
		MaxIdle: 50,
		//连接池中同时有多少条连接可以同时在连接
		MaxActive: 30,
		//连接超时时间
		IdleTimeout: 300 * time.Second,
		//具体配置redis连接
		Dial: func() (redis.Conn, error) {
			//1.打开连接
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				return nil, err
			}
			//2.访问认证
			if _, err = c.Do("AUTH"); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		//定时检查redis是否可用
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

//初始化连接池
func init() {
	pool = newRedisPool()
}

//向外暴露连接池
func RedisPool() *redis.Pool {
	return pool
}
