package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/lewgun/web-seed/pkg/errutil"

)

const (
	//ModeDebug run as debug
	ModeDebug = "debug"

	//ModeRelease run as release
	ModeRelease = "release"
)


type DataBase struct {
	DB       string `json:"dbname"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DSN      string `json:"-"`
}

type WebServer struct {
	HTTPPort int    `json:"port"`
	RunMode  string `json:"run_mode"`
	Prefix   string `json:"prefix"`
}

type Config struct {
	DataBase  `json:"database"`
	WebServer `json:"webserver"`
}

func (c *Config) init(path string) error {
	var err error
	if err = c.parse(path); err != nil {
		return fmt.Errorf("Can't load config from: %s with error: %v ", path, err)
	}

	if err = c.adjust(); err != nil {
		return fmt.Errorf("Adjust config failed.")
	}

	return c.check()
}

func (c *Config) adjust() error {

	c.RunMode = strings.ToLower(c.RunMode)

	if c.RunMode != ModeRelease {
		c.RunMode = ModeDebug
	}

	c.Prefix = strings.TrimSuffix(c.Prefix, "/")
	if c.Prefix[0] != '/' {
		c.Prefix = "/" + c.Prefix
	}

	//"root:123new@tcp(125.64.93.75:3306)/oss?charset=utf8"
	c.DSN = c.User + ":" + c.Password + "@tcp(" + c.DataBase.IP + ":" + strconv.Itoa(c.DataBase.Port) + ")/" + c.DB + "?charset=utf8&parseTime=true&loc=Local"

	return nil
}

func (c *Config) parse(path string) error {
	if path == "" {
		return errutil.ErrIllegalParam
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, c)
	fmt.Println(c)
	return err
}

//check检测配置参数是否完备
func (c *Config) check() error {
	return nil
}

//Load load一个配置
func Load(path string) (*Config, error) {
	c := &Config{}
	err := c.init(path)

	if err != nil {
		return nil, err
	}
	return c, nil
}
