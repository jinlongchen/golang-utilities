package cache

import (
    "log"
    "testing"
    "time"
)

func TestRedisCache_Get(t *testing.T) {
    var c Cache
    c = NewRedisCache(
        map[string]string{
            "c1": "127.0.0.1:6379",
        },
        "",
    )
    defer c.Close()
    var g []byte
    err := c.Get("test_cache", &g)
    if err != nil {
        log.Fatalln("get err:", err.Error())
    }
    log.Println("val:", g)
}
func TestRedisCache_Set(t *testing.T) {
    var c Cache
    c = NewRedisCache(
        map[string]string{
            "c1": "127.0.0.1:6379",
        },
        "",
    )
    defer c.Close()
    err := c.Set("test_cache", []byte(time.Now().String()), time.Hour)
    if err != nil {
        log.Fatalln("set err:", err.Error())
    }
}
