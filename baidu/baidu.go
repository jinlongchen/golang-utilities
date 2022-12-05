/*
 * Copyright (c) 2020. Jinlong Chen.
 */

package baidu

import (
    "fmt"
    "sync"

    "github.com/jinlongchen/golang-utilities/cache"
    "github.com/jinlongchen/golang-utilities/config"
    "github.com/jinlongchen/golang-utilities/log"
)

type Baidu struct {
    cache   cache.Cache
    memory  *sync.Map
    config  *config.Config
    logFunc func(string)
    quit    chan struct{}
}

func NewBaidu(cah cache.Cache, config *config.Config, logFunc func(string)) *Baidu {
    ret := &Baidu{
        cache:   cah,
        memory:  new(sync.Map),
        config:  config,
        logFunc: logFunc,
        quit:    make(chan struct{}),
    }

    return ret
}

func (bd *Baidu) Exit() error {
    close(bd.quit)
    return nil
}

func (bd *Baidu) logf(format string, args ...interface{}) {
    if bd.logFunc != nil {
        bd.logFunc(fmt.Sprintf(format, args...))
    }
    log.Infof(format, args...)
}
