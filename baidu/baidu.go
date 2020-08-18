/*
 * Copyright (c) 2020. Jinlong Chen.
 */

package baidu

import (
	"github.com/jinlongchen/golang-utilities/cache"
	"github.com/jinlongchen/golang-utilities/config"
)

type Baidu struct {
	cache  cache.Cache
	config *config.Config
	quit   chan struct{}
}

func NewBaidu(cah cache.Cache, config *config.Config) *Baidu {
	ret := &Baidu{
		cache:  cah,
		config: config,
		quit:   make(chan struct{}),
	}
	go func() {
		ret.fetchAccessTokensLoop()
	}()
	return ret
}

func (bd *Baidu) Exit() error {
	close(bd.quit)
	return nil
}
