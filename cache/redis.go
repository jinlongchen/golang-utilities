package cache

import (
    "context"
    "time"

    redisCache "github.com/go-redis/cache/v8"
    "github.com/go-redis/redis/v8"
    // redisCache "github.com/jinlongchen/redis-cache-go"
)

type RedisCache struct {
    rCache *redisCache.Cache
    ring   *redis.Ring
}

func NewRedisCache(addrs map[string]string, pwd string) Cache {
    ring := redis.NewRing(&redis.RingOptions{
        Addrs:    addrs,
        Password: pwd,
    })

    // codec := &redisCache.Codec{
    //     Redis: ring,
    //     Marshal: func(v interface{}) ([]byte, error) {
    //         return json.Marshal(v)
    //     },
    //     Unmarshal: func(b []byte, v interface{}) error {
    //         err := json.Unmarshal(b, v)
    //         if err != nil {
    //             log.Errorf(
    //                 `cannot unmarshal data: %v, %v`,
    //                 err,
    //                 hex.EncodeToString(b))
    //         }
    //         return err
    //     },
    // }
    // return &RedisCache{
    //     codec: codec,
    //     ring:  ring,
    // }

    rCache := redisCache.New(&redisCache.Options{
        Redis:      ring,
        LocalCache: redisCache.NewTinyLFU(1000, time.Minute),
    })

    return &RedisCache{
        rCache: rCache,
        ring:   ring,
    }
}

func (c *RedisCache) Delete(key string) error {
    return c.rCache.Delete(context.TODO(), key)
}

func (c *RedisCache) Get(key string, obj interface{}) error {
    err := c.rCache.Get(context.TODO(), key, &obj)
    if err == nil {
        return nil
    }
    return err
}

func (c *RedisCache) Set(key string, obj interface{}, timeout time.Duration) error {
    return c.rCache.Set(&redisCache.Item{
        Key:   key,
        Value: obj,
        TTL:   timeout,
    })
}
func (c *RedisCache) Close() error {
    return c.ring.Close()
}
