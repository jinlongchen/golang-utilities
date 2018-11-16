package aliapi

import "github.com/jinlongchen/golang-utilities/config"

type AliApiHelper struct {
	cfg *config.Config
}

func NewSAliApiHelper(cfg *config.Config) *AliApiHelper {
	return &AliApiHelper{cfg: cfg}
}
