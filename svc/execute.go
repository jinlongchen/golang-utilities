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

func (s *Executor) Init(env go_svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}

	cfgName := flag.String("conf", "./conf-file.json", "")
	aesKeyKey := flag.String("aesKeyKey", "jin^_^longchen", "")

	flag.Parse()

	s.cfg = config.NewConfig(*cfgName, *aesKeyKey)

	log.InitLogger(s.cfg.GetString("application.name"),
		log.LogLevel(s.cfg.GetString("log.level")),
		log.LOG_FORMAT_JSON,
		true)

	return nil
}

func (e *Executor) Start() error {
	banner.Print(e.s.GetName())
	e.s.Main(e.cfg)
	return nil
}

func (e *Executor) Stop() error {
	if e.s != nil {
		e.s.Exit()
	}
	return nil
}
