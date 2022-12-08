package cache

import (
	"context"
	"time"

	"github.com/rueian/rueidis"

	"github.com/jinlongchen/golang-utilities/json"
)

type RedisCache struct {
	rueidisCli rueidis.Client
}

func NewRedisCache(addrs []string, pwd string) Cache {
	redisC, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: addrs,
		Password:    pwd,
	})

	if err != nil {
		return nil
	}

	return &RedisCache{
		rueidisCli: redisC,
	}
}

func (c *RedisCache) Delete(key string) error {
	return c.rueidisCli.Do(context.Background(), c.rueidisCli.B().Del().Key(key).Build()).Error()
}

func (c *RedisCache) Get(key string, obj interface{}) error {
	return c.rueidisCli.Do(context.Background(), c.rueidisCli.B().Get().Key(key).Build()).DecodeJSON(obj)
}

func (c *RedisCache) Set(key string, obj interface{}, timeout time.Duration) error {
	data := json.ShouldMarshal(obj)
	v := c.rueidisCli.B().Set().Key(key).Value(string(data))
	if timeout > time.Duration(0) {
		return c.rueidisCli.Do(context.Background(), v.ExSeconds(int64(timeout.Seconds())).Build()).Error()
	}
	return c.rueidisCli.Do(context.Background(), v.Build()).Error()
}

func (c *RedisCache) Close() error {
	c.rueidisCli.Close()
	return nil
}
