package wechat

import (
	"github.com/jinlongchen/golang-utilities/cache"
	"github.com/jinlongchen/golang-utilities/config"
)

type Wechat struct {
	cache  cache.Cache
	config *config.Config
}

func NewWechat(cah cache.Cache, config *config.Config) *Wechat {
	ret := &Wechat{cache: cah, config: config}
	return ret
}
