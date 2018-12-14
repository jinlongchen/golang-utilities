package svc

import (
	go_svc "github.com/judwhite/go-svc/svc"
	"path/filepath"
	"os"
	"syscall"
	"flag"
	"github.com/jinlongchen/golang-utilities/config"
	"github.com/jinlongchen/golang-utilities/log"
	"github.com/jinlongchen/golang-utilities/banner"
)

type Executor struct {
	s   Service
	cfg *config.Config
}

func Execute(s Service) {
	e := &Executor{}
	e.s = s
	if err := go_svc.Run(e, syscall.SIGINT, syscall.SIGTERM); err != nil {
		panic(err)
	}
}

func (e *Executor) Init(env go_svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}

	cfgName := flag.String("conf", "./conf-file.toml", "")

	flag.Parse()

	e.cfg = config.NewConfig(*cfgName)

	log.InitLogger(e.cfg.GetString("application.name"),
		log.LogLevel(e.cfg.GetString("log.level")),
		log.LOG_FORMAT_JSON,
		true)

	log.Infof("log level:%s", e.cfg.GetString("log.level"))
	return nil
}

func (e *Executor) Start() error {
	e.s.Main(e.cfg)
	banner.Print(e.s.GetName())
	return nil
}

func (e *Executor) Stop() error {
	if e.s != nil {
		e.s.Exit()
	}
	return nil
}
