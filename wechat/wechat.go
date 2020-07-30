package wechat

import (
	"github.com/jinlongchen/golang-utilities/cache"
	"github.com/jinlongchen/golang-utilities/config"
)

type Wechat struct {
	cache  cache.Cache
	config *config.Config
	quit   chan struct{}
}

func NewWechat(cah cache.Cache, config *config.Config) *Wechat {
	ret := &Wechat{
		cache: cah,
		config: config,
		quit: make(chan struct{}),
	}
	go func() {
		ret.fetchAccessTokensLoop()
	}()
	return ret
}
