package main

import (
	"flag"
	"github.com/yylt/example-swagger/api"
	"github.com/yylt/example-swagger/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	gsd "github.com/yylt/gorilla-session-django"
)

func main() {
	//address options
	addr := flag.String("addr", ":8089", "listen address")
	basepath := flag.String("baseurl", "/api/v1", "base url prefix")
	memcacheAddr := flag.String("memcache", "memcached.openstack.svc.cluster.local:11211", "memcached address")
	flag.Parse()

	memcli, err := gsd.NewMemCli(&gsd.MemCfg{
		Endpoint: *memcacheAddr,
		User:     "",
		Password: "",
	})
	if err != nil {
		panic(err)
	}
	logger := log.NewLogfmtLogger(os.Stdout)
	//api init, include wrap handler
	a, err := api.New(logger, *basepath)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	engine.Use(
		middleware.SessionAuthMid(memcli),
		gin.WrapH(a.Handler),
	)

	err = engine.Run(*addr)
	if err != nil {
		panic(err)
	}
}
