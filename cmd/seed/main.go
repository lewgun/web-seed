package main

import (
	"flag"
	"fmt"
	"time"

	//config
	"github.com/lewgun/web-seed/pkg/config"

	//model
	"github.com/lewgun/web-seed/pkg/model"

	//logger
	"github.com/lewgun/web-seed/pkg/zlog"
	_ "github.com/lewgun/web-seed/pkg/zlog/logrus"

	//webserver
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	//sessions
	"github.com/gin-gonic/contrib/sessions"
	mySession "github.com/lewgun/web-seed/pkg/session/mysql"

	//controller
	"github.com/lewgun/web-seed/pkg/controller"

	//logger
	"github.com/Sirupsen/logrus"

)

var (
	confPath = flag.String("conf", "./config.json", "the path to the config file")
)

func main() {
	c, store := mustPrepare()
	powerOn(c, store)
	powerOff()
}

func setupRouter(r *gin.Engine) {

	r.Use(static.Serve("/", static.LocalFile("web", false)))

	r.StaticFile("/", "web/index.html")

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	controller.New(r).SetupRouters()

}

func powerOn(c *config.Config, store mySession.MySQLStore) {

	if c.RunMode == config.ModeRelease {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	rLogger := zlog.L.RealLogger(zlog.DriverLogrus)
	if rLogger != nil {
		if l, ok := rLogger.(*logrus.Logger); ok {
			r.Use(ginrus.Ginrus(l, time.RFC3339, false))
		}
	}

	r.Use(sessions.Sessions("session", store))

	r.Use(cors.Default())

	setupRouter(r)

	fmt.Println("passport is running at: ", c.HTTPPort)

	r.Run(fmt.Sprintf(":%d", c.HTTPPort)) // listen and serve on 0.0.0.0:8080
}

func powerOff() {
	zlog.L.PowerOff()

}

func mustPrepare() (*config.Config, mySession.MySQLStore) {

	//configure
	c, err := config.Load(*confPath)
	if err != nil {
		panic(err)
	}

	//logger
	err = zlog.BootUp()
	if err != nil {
		panic(err)
	}

	// model
	m := model.SharedInstance(c.DSN, 100)
	if m == nil {
		panic("can't boot up storage.")
	}

	store, err := mySession.NewMySQLStore(c.DSN, 3600, "something very secret")
	if err != nil {
		panic(err)
	}

	return c, store

}
