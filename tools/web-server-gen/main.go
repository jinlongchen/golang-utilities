package main

import (
	//"fmt"
	//"log"
	//"io/ioutil"
	"os"

	"github.com/urfave/cli"
	"runtime"
	"path/filepath"
	"strings"
	"log"
	"fmt"
	"io/ioutil"
)

var version = "dev"

func main() {
	app := cli.NewApp()
	app.Name = "web-server-gen"
	app.Usage = "web server startup generator."
	app.Version = version
	app.Author = "Jinlong Chen"
	app.Commands = []cli.Command{
		{
			Name:      "generate",
			ArgsUsage: "",
			Aliases:   []string{"g"},
			Usage:     "Generates startup",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output,o",
					Value: "./",
					Usage: "Output path",
				},
				cli.StringFlag{
					Name:  "name,n",
					Value: "",
					Usage: "web server name",
				},
			},
			Action: func(c *cli.Context) {
				var err error

				outputPath := c.String("output")
				if runtime.GOOS != "windows" {
					if strings.HasPrefix(outputPath, "~") {
						outputPath = expandUser(outputPath)
					} else if !strings.HasPrefix(outputPath, "/") {
						wd, _ := os.Getwd()
						outputPath = filepath.Join(wd, outputPath)
					}
				}

				webServerName := c.String("name")

				err = os.MkdirAll(fmt.Sprintf("%s/app/%s", outputPath, webServerName), 0777)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = os.MkdirAll(fmt.Sprintf("%s/%s/helper", outputPath, webServerName), 0777)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = os.MkdirAll(fmt.Sprintf("%s/%s/biz/handler/test", outputPath, webServerName), 0777)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = os.MkdirAll(fmt.Sprintf("%s/%s/biz/data/test", outputPath, webServerName), 0777)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = os.MkdirAll(fmt.Sprintf("%s/%s/biz/model", outputPath, webServerName), 0777)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = os.MkdirAll(fmt.Sprintf("%s/%s/biz/error-codes", outputPath, webServerName), 0777)
				if err != nil {
					log.Fatalln(err.Error())
				}

				srcDir := filepath.Join(os.Getenv("GOPATH"), "src")
				pkgPath, _ := filepath.Rel(srcDir, outputPath)

				err = ioutil.WriteFile(
					fmt.Sprintf("%s/app/%s/%s.go", outputPath, webServerName, webServerName),
					getAppFileContent(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = ioutil.WriteFile(
					fmt.Sprintf("%s/%s/%s.go", outputPath, webServerName, webServerName),
					getPkgFile1(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = ioutil.WriteFile(
					fmt.Sprintf("%s/%s/version.go", outputPath, webServerName),
					getPkgFile2(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = ioutil.WriteFile(
					fmt.Sprintf("%s/%s/wait.go", outputPath, webServerName),
					getPkgFile3(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = ioutil.WriteFile(
					fmt.Sprintf("%s/%s/http.go", outputPath, webServerName),
					getPkgFile4(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = ioutil.WriteFile(
					fmt.Sprintf("%s/%s/helper/context.go", outputPath, webServerName),
					getPkgFile5(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = ioutil.WriteFile(
					fmt.Sprintf("%s/%s/helper/http.go", outputPath, webServerName),
					getPkgFile11(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = ioutil.WriteFile(
					fmt.Sprintf("%s/%s/biz/data/test/manager.go", outputPath, webServerName),
					getPkgFile6(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = ioutil.WriteFile(
					fmt.Sprintf("%s/%s/biz/data/test/manager_sql.go", outputPath, webServerName),
					getPkgFile7(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = ioutil.WriteFile(
					fmt.Sprintf("%s/%s/biz/model/test.go", outputPath, webServerName),
					getPkgFile8(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = ioutil.WriteFile(
					fmt.Sprintf("%s/%s/biz/handler/test/test.go", outputPath, webServerName),
					getPkgFile9(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}
				err = ioutil.WriteFile(
					fmt.Sprintf("%s/%s/biz/error-codes/error-codes.go", outputPath, webServerName),
					getPkgFile10(pkgPath, webServerName),
					0666)
				if err != nil {
					log.Fatalln(err.Error())
				}

			},
		},
	}
	app.Run(os.Args)
}

func expandUser(s string) string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	home := os.Getenv(env)
	if home == "" {
		return s
	}

	if len(s) >= 2 && s[0] == '~' && os.IsPathSeparator(s[1]) {
		if runtime.GOOS == "windows" {
			s = filepath.ToSlash(filepath.Join(home, s[2:]))
		} else {
			s = filepath.Join(home, s[2:])
		}
	}
	return os.Expand(s, func(env string) string {
		if env == "HOME" {
			return home
		}
		return os.Getenv(env)
	})
}

func getAppFileContent(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package main

import (
	"github.com/jinlongchen/golang-utilities/svc"
	"%s/%s"
)

func main() {
	svc.Execute(%s.NewService())
}
`, pkgPath, webServerName, webServerName))
}

func getPkgFile1(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package %s

import (
	"fmt"
	"net/url"
	"time"

	"github.com/jinlongchen/golang-utilities/cache"
	"github.com/jinlongchen/golang-utilities/config"
	"github.com/jinlongchen/golang-utilities/database"
	"github.com/jinlongchen/golang-utilities/log"
	"github.com/jinlongchen/golang-utilities/version"
	"%s/%s/helper"
)

type Service struct {
	startTime   time.Time
	ctx         *helper.Context
	httpHandler *HttpHandler
	waitGroup   WaitGroupWrapper
}

func NewService() *Service {
	return &Service{
		startTime: time.Now(),
	}
}

func (d *Service) GetName() string {
	return "%s"
}

func (d *Service) Main(cfg *config.Config) {
	fmt.Println(version.String(cfg.GetString("application.name"), BINARY))

	connStr, err := url.Parse(cfg.GetString("database.pg.connStr"))
	if err != nil {
		log.Fatalf("conn str err %%s", err.Error())
	}

	connection := &database.SQLConnection{
		URL: connStr,
	}

	cacheServers := cfg.GetStringMapString("cache.redis.addresses")
	redisCache := cache.NewRedisCache(
		cacheServers,
		cfg.GetString("cache.redis.password"),
	)

	ctx := &helper.Context{
		Config:   cfg,
		Database: connection.GetDatabase(),
		Cache:    redisCache,
	}

	httpHandler := newHttpHandler(ctx)

	d.waitGroup.Wrap(func() {
		httpHandler.Serve()
	})

	d.httpHandler = httpHandler
}

func (d *Service) Exit() {
	d.httpHandler.Exit()
	d.waitGroup.Wait()
}
`, webServerName, pkgPath, webServerName, webServerName))
}

func getPkgFile2(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package %s

const (
	BINARY = "1.0"
)
`, webServerName))
}

func getPkgFile3(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package %s

import (
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}
`, webServerName))
}

func getPkgFile4(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package %s
import (
	"context"

	"github.com/jinlongchen/golang-utilities/converter"
	"github.com/jinlongchen/golang-utilities/crypto"
	"github.com/jinlongchen/golang-utilities/log"
	"github.com/jinlongchen/golang-utilities/version"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/gorilla/sessions"
	"%s/%s/helper"
	test_data "%s/%s/biz/data/test"
	test_handler "%s/%s/biz/handler/test"
)

type HttpHandler struct {
	echoEngine *echo.Echo
	ctx        *helper.Context
	quit       chan struct{}
}

func newHttpHandler(ctx *helper.Context) *HttpHandler {
	return &HttpHandler{
		ctx:  ctx,
		quit: make(chan struct{}),
	}
}

func (httpH *HttpHandler) initRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: " ${status}，${method}:${uri}\n",
	}))
	e.HideBanner = true
	e.HidePort = true
	e.Logger.SetLevel(1)
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 9,
	}))

	origins := converter.AsStringArray(httpH.ctx.Config.GetValue("http.header.origins"))
	if origins != nil {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     origins,
			AllowCredentials: true,
			ExposeHeaders:    []string{"X-Total-Count"},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlExposeHeaders, echo.HeaderXCSRFToken, "X-Total-Count"},
			AllowMethods:     []string{echo.POST, echo.GET, echo.OPTIONS},
		}))
	}

	e.Use(session.Middleware(sessions.NewCookieStore(crypto.String(httpH.ctx.Config.GetString("application.name")).GetMd5())))

	if httpH.ctx.Config.GetBool("http.secure.csrf") {
		e.Use(middleware.CSRF())
	}


	test_handler.NewHandler(httpH.ctx,
		&test_data.SQLManager{DB: httpH.ctx.Database}).InitWith(e)

	e.GET("/version", func(ctx echo.Context) error {
		return ctx.String(200, version.String(httpH.ctx.Config.GetString("application.name"), BINARY))
	})

	return e
}

func (httpH *HttpHandler) Serve() {
	httpH.echoEngine = httpH.initRouter()

	log.Infof("http server started on %%s", httpH.ctx.Config.GetString("network.http.listenOn"))
	err := httpH.echoEngine.Start(httpH.ctx.Config.GetString("network.http.listenOn"))

	log.Fatalf("http server err %%s", err.Error())
}

func (httpH *HttpHandler) Exit() {
	close(httpH.quit)
	httpH.echoEngine.Shutdown(context.Background())
}

`, webServerName, pkgPath, webServerName, pkgPath, webServerName, pkgPath, webServerName))
}

func getPkgFile5(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package helper

import (
	"github.com/jinlongchen/golang-utilities/cache"
	"github.com/jinlongchen/golang-utilities/config"
	"github.com/jmoiron/sqlx"
)

type Context struct {
	Config   *config.Config
	Database *sqlx.DB
	Cache    cache.Cache
}

`))
}

func getPkgFile6(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package test

import (
	"%s/%s/biz/model"
)

type Manager interface {
	Storage
}

type Storage interface {
	GetTestItems() ([]*model.TestItem, int, error)
	GetTestItem(id string) (testItem *model.TestItem, err error)
}

`, pkgPath, webServerName))
}

func getPkgFile7(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package test

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"
	"github.com/jinlongchen/golang-utilities/database/pg"
	"github.com/jinlongchen/golang-utilities/log"
	"github.com/jinlongchen/golang-utilities/errors"
	"strings"
	"database/sql"
	"%s/%s/biz/model"
)

const (
	SQL_DB_NAME = "xx_test"
)

var migrations = &migrate.MemoryMigrationSource{
	Migrations: []*migrate.Migration{
		{
			Id: "1",
			Up: []string{
				` + "`" + `CREATE TABLE IF NOT EXISTS ` + "`" + ` + SQL_DB_NAME + ` + "`" + ` (
		id			varchar(255) NOT NULL PRIMARY KEY,
		image		jsonb NULL,
		url 		text NULL
	)` + "`" + `,
			},
			Down: []string{
				` + "`" + `DROP TABLE ` + "`" + ` + SQL_DB_NAME + ` + "`" + `` + "`" + `,
			},
		},
	},
}

type SQLManager struct {
	DB *sqlx.DB
}

type sqlData struct {
	ID    string       ` + "`" + `db:"id"` + "`" + `
	Image *pg.JsonbMap ` + "`" + `db:"image"` + "`" + `
	URL   string       ` + "`" + `db:"url"` + "`" + `
}

var sqlParams = []string{
	"id",
	"image",
	"url",
}

func sqlDataFromTestItem(d *model.TestItem) *sqlData {
	image := &pg.JsonbMap{}
	image.From(d.Image)
	return &sqlData{
		ID:    d.ID,
		Image: image,
		URL:   d.URL,
	}
}

func (d *sqlData) ToTestItem() *model.TestItem {
	var err error
	ret := &model.TestItem{
		ID: d.ID,
		Image: &model.Image{},
		URL:   d.URL,
	}
	err = d.Image.To(ret.Image)
	if err != nil {
		log.Debugf("Image err:%s", err.Error())
	}

	return ret
}

func (m *SQLManager) CreateSchemas() (int, error) {
	if m.DB == nil {
		return 0, errors.New("db error")
	}
	migrate.SetTable(` + "`" + `` + "`" + ` + SQL_DB_NAME + ` + "`" + `_migration` + "`" + `)
	n, err := migrate.Exec(m.DB.DB, m.DB.DriverName(), migrations, migrate.Up)
	if err != nil {
		return 0, fmt.Errorf("could not migrate sql schema, applied %d migrations", n)
	}
	return n, nil
}

func (m *SQLManager) GetTestItems() (testItems []*model.TestItem, totalCount int, err error) {
	var d = make([]sqlData, 0)
	testItems = make([]*model.TestItem, 0)

	if m.DB == nil {
		return nil, 0, errors.New("db error")
	}

	err = m.DB.Get(
		&totalCount,
		m.DB.Rebind(` + "`" + `SELECT COUNT(id) FROM ` + "`" + ` + SQL_DB_NAME + ` + "`" + ` ` + "`" + `))

	if err == nil {
		err = m.DB.Select(
			&d,
			m.DB.Rebind(` + "`" + `SELECT ` + "`" + `+strings.Join(sqlParams, ", ")+` + "`" + ` FROM ` + "`" + ` + SQL_DB_NAME + ` + "`" + `  ORDER BY id DESC ` + "`" + `))
		if err != nil {
			log.Errorf("get testItems err:%s", err.Error())
		}
	} else {
		log.Errorf("get count err:%s", err.Error())
	}
	if err == sql.ErrNoRows {
		return nil, 0, err
	} else if err != nil {
		return nil, 0, err
	}

	for _, k := range d {
		testItems = append(testItems, k.ToTestItem())
	}
	return testItems, totalCount, nil
}

func (m *SQLManager) GetTestItem(id string) (testItem *model.TestItem, err error) {
	var d sqlData

	if m.DB == nil {
		return &model.TestItem{}, errors.New("db error")
	}

	if err := m.DB.Get(&d, m.DB.Rebind(` + "`" + `SELECT ` + "`" + `+strings.Join(sqlParams, ", ")+` + "`" + ` FROM ` + "`" + ` + SQL_DB_NAME + ` + "`" + ` WHERE id=?` + "`" + `), id); err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	r := d.ToTestItem()
	if err != nil {
		return nil, err //errors.WithStack(err)
	}
	return r, nil
}

`, pkgPath, webServerName))
}

func getPkgFile8(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package model

type Image struct {
	ID          string  ` + "`" + `json:"id" db:"id"` + "`" + `
	URL         string  ` + "`" + `json:"url" db:"url"` + "`" + `
	Width       int     ` + "`" + `json:"width" db:"width"` + "`" + `
	AspectRatio float32 ` + "`" + `json:"aspect_ratio" db:"aspect_ratio"` + "`" + `
}

type TestItem struct {
	ID    string ` + "`" + `json:"id" db:"id"` + "`" + `
	Image *Image ` + "`" + `json:"image" db:"image"` + "`" + `
	URL   string ` + "`" + `json:"url" db:"url"` + "`" + `
}
`))
}

func getPkgFile9(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package test

import (
	"github.com/labstack/echo"
	test_data "%s/%s/biz/data/test"
	"%s/%s/biz/error-codes"
	"%s/%s/helper"
	"net/http"
)

type Handler struct {
	ctx     *helper.Context
	manager test_data.Manager
}

func NewHandler(ctx *helper.Context, manager test_data.Manager) *Handler {
	return &Handler{
		ctx:     ctx,
		manager: manager,
	}
}

func (httpH *Handler) InitWith(e *echo.Echo) {
	e.GET("/test_items", httpH.GetBanners)
}

func (httpH *Handler) InitWithGroup(g *echo.Group) {
	g.GET("/test_items", httpH.GetBanners)
}

func (httpH *Handler) GetBanners(ctx echo.Context) error {
	//test only
	testItems, _, err := httpH.manager.GetTestItems()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, error_codes.NewErrorCode(error_codes.ErrorCodeServerError))
	}

	return helper.ResponseItems(ctx, testItems, len(testItems), 0, nil)
}
`, pkgPath, webServerName,
pkgPath, webServerName,
pkgPath, webServerName))
}

func getPkgFile10(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package error_codes

import "github.com/jinlongchen/golang-utilities/errors"

type ErrorCode string

const (
	ErrorCodeNone           ErrorCode = ""
	ErrorCodeUnknown        ErrorCode = "Unknown"
	ErrorCodeBadRequest     ErrorCode = "BadRequest"
	ErrorCodeUnauthorized   ErrorCode = "Unauthorized"
	ErrorCodeServerError    ErrorCode = "ServerError"
)

var (
	ErrorCodeDescMap = map[ErrorCode]string{
		ErrorCodeNone:           "",
		ErrorCodeUnknown:        "未知错误",
		ErrorCodeBadRequest:     "请求数据错误，请检查你的输入",
		ErrorCodeUnauthorized:   "请登录后进行操作",
		ErrorCodeServerError:    "服务器错误",
	}
)

func SetErrorCode(object map[string]interface{}, code ErrorCode) map[string]interface{} {
	if object == nil {
		object = make(map[string]interface{})
	}
	object["err_code"] = string(code)
	object["err_msg"] = ErrorCodeDescMap[code]
	return object
}

func NewErrorCode(code ErrorCode) map[string]interface{} {
	object := make(map[string]interface{})

	object["err_code"] = string(code)
	object["err_msg"] = ErrorCodeDescMap[code]

	return object
}

func SetErrorCodeFromError(object map[string]interface{}, err error) map[string]interface{} {
	if object == nil {
		return NewErrorCode(ErrorCodeNone)
	}

	if err == nil {
		return SetErrorCode(object, ErrorCodeNone)
	}
	if withCodeErr, ok := err.(*errors.WithCodeError); ok {
		object["err_code"] = withCodeErr.Code()
		object["err_msg"] = withCodeErr.Error()
		return object
	}

	return SetErrorCode(object, ErrorCodeUnknown)
}

func NewErrorCodeFromError(err error) map[string]interface{} {
	if err == nil {
		return NewErrorCode(ErrorCodeNone)
	}

	if withCodeErr, ok := err.(*errors.WithCodeError); ok {
		object := make(map[string]interface{})
		object["err_code"] = withCodeErr.Code()
		object["err_msg"] = withCodeErr.Error()
		return object
	}

	return NewErrorCode(ErrorCodeUnknown)
}

`))
}

func getPkgFile11(pkgPath string, webServerName string) []byte {
	return []byte(fmt.Sprintf(`package helper

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"%s/%s/biz/error-codes"
)

func ResponseItems(ctx echo.Context, data interface{}, totalCount int, pageIndex int, err error) error {
	result := map[string]interface{}{
		"data":        data,
		"total_count": totalCount,
		"page_index":  pageIndex,
	}
	if err != nil {
		log.Debug(err)
	}

	//	ctx.Response().Header().Set("X-Total-Count", fmt.Sprintf("%d",totalCount))

	return ctx.JSON(http.StatusOK,
		error_codes.SetErrorCodeFromError(result, err),
	)
}

func ResponseItem(ctx echo.Context, data interface{}, err error) error {
	result := map[string]interface{}{
		"data":        data,
	}
	if err != nil {
		log.Debug(err)
	}

	return ctx.JSON(http.StatusOK,
		error_codes.SetErrorCodeFromError(result, err),
	)
}

`, pkgPath, webServerName))
}
