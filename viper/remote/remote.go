// Copyright © 2015 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package remote integrates the remote features of Viper.
package remote

import (
	"bytes"
	"io"
	"os"

	"github.com/jinlongchen/golang-utilities/viper"
	crypt "github.com/jinlongchen/golang-utilities/viper/crypt/config"
)

type remoteConfigProvider struct{}

func (rc remoteConfigProvider) Get(rp viper.RemoteProvider) (io.Reader, error) {
	return rc.fetchConfig(rp)
}

func (rc remoteConfigProvider) Watch(rp viper.RemoteProvider) (io.Reader, error) {
	return rc.fetchConfig(rp)
}

func (rc remoteConfigProvider) fetchConfig(rp viper.RemoteProvider) (io.Reader, error) {
	cm, err := getConfigManager(rp)
	if err != nil {
		return nil, err
	}
	b, err := cm.Get(rp.Path())
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (rc remoteConfigProvider) WatchChannel(rp viper.RemoteProvider) (<-chan *viper.RemoteResponse, chan bool) {
	cm, err := getConfigManager(rp)
	if err != nil {
		return nil, nil
	}
	quit := make(chan bool)
	quitwc := make(chan bool)
	viperResponsCh := make(chan *viper.RemoteResponse)
	cryptoResponseCh := cm.Watch(rp.Path(), quit)

	go rc.convertResponse(cryptoResponseCh, viperResponsCh, quitwc, quit)

	return viperResponsCh, quitwc
}

func (rc remoteConfigProvider) convertResponse(cr <-chan *crypt.Response, vr chan<- *viper.RemoteResponse, quitwc <-chan bool, quit chan<- bool) {
	for {
		select {
		case <-quitwc:
			quit <- true
			return
		case resp := <-cr:
			vr <- &viper.RemoteResponse{
				Error: resp.Error,
				Value: resp.Value,
			}
		}
	}
}

func getConfigManager(rp viper.RemoteProvider) (crypt.ConfigManager, error) {
	var cm crypt.ConfigManager
	var err error

	if rp.SecretKeyring() != "" {
		kr, err := os.Open(rp.SecretKeyring())
		if err != nil {
			return nil, err
		}
		defer kr.Close()

		cm, err = newConfigManagerWithKeyring(rp, kr)
	} else {
		cm, err = newStandardConfigManager(rp)
	}

	if err != nil {
		return nil, err
	}
	return cm, nil
}

func newConfigManagerWithKeyring(rp viper.RemoteProvider, kr *os.File) (crypt.ConfigManager, error) {
	if rp.Provider() == "etcd" {
		return crypt.NewEtcdConfigManager([]string{rp.Endpoint()}, kr)
	}
	return crypt.NewConsulConfigManager([]string{rp.Endpoint()}, kr)
}

func newStandardConfigManager(rp viper.RemoteProvider) (crypt.ConfigManager, error) {
	if rp.Provider() == "etcd" {
		return crypt.NewStandardEtcdConfigManager([]string{rp.Endpoint()})
	}
	return crypt.NewStandardConsulConfigManager([]string{rp.Endpoint()})
}

func init() {
	viper.RemoteConfig = &remoteConfigProvider{}
}
