package main

import (
	"encoding/json"
	"flag"
	"github.com/crusj/laravelAdminCmd/cmd"
	_ "github.com/crusj/laravelAdminCmd/init"
	logger "github.com/crusj/logger"
	"io/ioutil"
	"os"
)

var (
	path   string
	make   cmd.Make = cmd.Make{}
	config Config
)

type Config struct {
	Make []cmd.MakeConfig `json:"make"`
}

func init() {
	flag.StringVar(&path, "path", "cmd.json", "配置文件路径")
	flag.Var(&make, "make", "php artisan admin:make")
	flag.Parse()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			logger.Info("执行操作失败:%v",err)
		}
	}()

	if path == "" {
		path = "cmd.json"
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Painc("配置文件%v,不存在", path)
	}
	if err := parseConfig(); err != nil {
		logger.Painc("解析配置文件%v失败,%v", path, err)
	}

	//make命令
	make.Config = config.Make
	make.CheckAndRunMake()

}
func parseConfig() error {
	if file, err := os.Open(path); err != nil {
		logger.Warn("无法打开配置文件%v,", err)
		return err
	} else {
		fileData, _ := ioutil.ReadAll(file)
		if err := json.Unmarshal(fileData, &config); err != nil {
			logger.Warn("读取配置文件%v失败,", path, err)
			return err
		}
	}
	return nil
}
