package database

import (
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/jinlongchen/golang-utilities/retry"
	"fmt"
	"github.com/jinlongchen/golang-utilities/log"
)

type SQLConnection struct {
	db      *sqlx.DB
	URL     *url.URL
}

func cleanURLQuery(c *url.URL) *url.URL {
	cleanurl := new(url.URL)
	*cleanurl = *c

	q := cleanurl.Query()
	q.Del("max_conns")
	q.Del("max_idle_conns")
	q.Del("max_conn_lifetime")

	cleanurl.RawQuery = q.Encode()
	return cleanurl
}

func (c *SQLConnection) GetDatabase() *sqlx.DB {
	if c.db != nil {
		return c.db
	}

	var err error
	clean := cleanURLQuery(c.URL)

	if err = retry.Run(time.Second*15, time.Minute, func() error {
		u := clean.String()
		if clean.Scheme == "mysql" {
			u = strings.Replace(u, "mysql://", "", -1)
		}

		if c.db, err = sqlx.Open(clean.Scheme, u); err != nil {
			return fmt.Errorf("could not connect to SQL: %s", err)
		} else if err := c.db.Ping(); err != nil {
			return fmt.Errorf("could not connect to SQL: %s", err)
		}

		return nil
	}); err != nil {
		log.Errorf("Could not Connect to SQL: %v", err)
		return nil
	}

	maxConns := maxParallelism() * 2
	if v := c.URL.Query().Get("max_conns"); v != "" {
		s, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Warnf( "max_conns value %s could not be parsed to int: %s", v, err)
		} else {
			maxConns = int(s)
		}
	}

	maxIdleConns := maxParallelism()
	if v := c.URL.Query().Get("max_idle_conns"); v != "" {
		s, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Warnf( "max_idle_conns value %s could not be parsed to int: %s", v, err)
		} else {
			maxIdleConns = int(s)
		}
	}

	maxConnLifetime := time.Duration(0)
	if v := c.URL.Query().Get("max_conn_lifetime"); v != "" {
		s, err := time.ParseDuration(v)
		if err != nil {
			log.Warnf( "max_conn_lifetime value %s could not be parsed to int: %s", v, err)
		} else {
			maxConnLifetime = s
		}
	}

	c.db.SetMaxOpenConns(maxConns)
	c.db.SetMaxIdleConns(maxIdleConns)
	c.db.SetConnMaxLifetime(maxConnLifetime)

	return c.db
}

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

