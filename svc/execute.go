package svc

import (
	"flag"
	"github.com/jinlongchen/golang-utilities/config"
	"github.com/jinlongchen/golang-utilities/log"
	go_svc "github.com/judwhite/go-svc/svc"
	"os"
	"path/filepath"
	"syscall"
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
		err := os.Chdir(dir)
		if err != nil {
			panic(err.Error())
		}
	}

	cfgName := flag.String("conf", "conf-file.toml", "")

	flag.Parse()

	e.cfg = config.NewConfig(*cfgName)

	log.Config(e.cfg.GetString("application.name"),
		log.Level(e.cfg.GetString("log.level")),
		e.cfg.GetBool("log.console"),
		e.cfg.GetString("log.filename"),
		e.cfg.GetInt("log.maxSize"),
		e.cfg.GetInt("log.maxBackups"),
		e.cfg.GetInt("log.maxAge"),
	)

	log.Infof("log level:%s", e.cfg.GetString("log.level"))
	return nil
}

func (e *Executor) Start() error {
	e.s.Main(e.cfg)
	return nil
}

func (e *Executor) Stop() error {
	if e.s != nil {
		e.s.Exit()
	}
	log.Flush()
	if e.cfg != nil {
		e.cfg.Exit()
	}
	return nil
}
