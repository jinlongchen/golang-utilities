package cache

import (
	redisCache "github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
	"time"
)

type RedisCache struct {
	codec *redisCache.Codec
	ring  *redis.Ring
}

func NewRedisCache(addrs map[string]string, pwd string) Cache {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs:    addrs,
		Password: pwd,
	})
	codec := &redisCache.Codec{
		Redis: ring,
		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
	return &RedisCache{
		codec: codec,
		ring:  ring,
	}
}

func (c *RedisCache) Delete(key string) error {
	return c.codec.Delete(key)
}

func (c *RedisCache) Get(key string, obj interface{}) error {
	err := c.codec.Get(key, &obj)
	if err == nil {
		return nil
	}
	return err
}

func (c *RedisCache) Set(key string, obj interface{}, timeout time.Duration) error {
	return c.codec.Set(&redisCache.Item{
		Key:        key,
		Object:     obj,
		Expiration: timeout,
	})
}
func (c *RedisCache) Close() error {
	return c.ring.Close()
}
