package aliapi

import "github.com/jinlongchen/golang-utilities/config"

type AliApiHelper struct {
    cfg *config.Config
}

func NewAliApiHelper(cfg *config.Config) *AliApiHelper {
    return &AliApiHelper{cfg: cfg}
}
