package svc

import (
	"flag"
	"net/url"
	"os"
	"path/filepath"
	"syscall"

	go_svc "github.com/judwhite/go-svc/svc"

	"github.com/jinlongchen/golang-utilities/config"
	"github.com/jinlongchen/golang-utilities/log"
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
	remoteConfigURL := os.Getenv("REMOTE_CONFIG_URL")
	localConfigURL := os.Getenv("LOCAL_CONFIG_URL")
	if localConfigURL == "" {
		cfgName := flag.String("conf", "conf-file.toml", "")
		flag.Parse()
		localConfigURL = *cfgName
	}

	if remoteConfigURL != "" {
		uRL, err := url.Parse(remoteConfigURL)
		if err == nil {
			log.Infof("Remote Config -> Schema: %s, Host: %s, Path: %s", uRL.Scheme, uRL.Host, uRL.Path)
			e.cfg = config.NewRemoteConfig(uRL.Scheme, uRL.Host, uRL.Path)
			if localConfigURL != "" {
				_ = e.cfg.Save(localConfigURL)
			}
		}
	}
	if e.cfg == nil {
		if localConfigURL != "" {
			e.cfg = config.NewConfig(localConfigURL)
		}
	}

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
