package svc

import (
	"github.com/jinlongchen/golang-utilities/config"
)

type Service interface {
	GetName() string
	Main(config *config.Config)
	Exit()
}
