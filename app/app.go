package app

import (
	"flag"
	"fmt"
	"ideal/module"
	"os"
	"os/exec"
	"path/filepath"
)

func newOptions(opts ...module.Option) module.Options {
	var confPath, logDir *string
	opt := module.Options{
		Debug: true,
		Parse: true,
	}

	for _, o := range opts {
		o(&opt)
	}

	if opt.Parse {
		confPath = flag.String("conf", "", "配置文件路径")
		logDir = flag.String("log", "", "日志文件目录")
		flag.Parse()
	}

	ApplicationDir := ""
	if opt.WorkDir != "" {
		_, err := os.Open(opt.WorkDir)
		if err != nil {
			panic(err)
		}
		os.Chdir(opt.WorkDir)
		ApplicationDir, err = os.Getwd()
	} else {
		var err error
		ApplicationDir, err = os.Getwd()
		if err != nil {
			file, _ := exec.LookPath(os.Args[0])
			ApplicationPath, _ := filepath.Abs(file)
			ApplicationDir, _ = filepath.Split(ApplicationPath)
		}
	}
	opt.WorkDir = ApplicationDir
	defaultConfPath := fmt.Sprintf("%s/bin/conf/server.json", ApplicationDir)
	defaultLogPath := fmt.Sprintf("%s/bin/logs", ApplicationDir)

	if opt.ConfPath == "" {
		if confPath == nil || *confPath == "" {
			opt.ConfPath = defaultConfPath
		} else {
			opt.ConfPath = *confPath
		}
	}

	if opt.LogDir == "" {
		if logDir == nil || *logDir == "" {
			opt.LogDir = defaultLogPath
		} else {
			opt.LogDir = *logDir
		}
	}

	if _, err := os.Open(opt.ConfPath); err != nil {
		panic(fmt.Sprintf("配置文件读取异常 %v", err))
	}

	_, err := os.Open(opt.LogDir)
	if err != nil {
		os.Mkdir(opt.LogDir, os.ModePerm)
	}

	return opt
}

func NewApp(opts ...module.Option) module.App {
	app := new(App)
	app.opts = newOptions(opts...)
	return app
}

type App struct {
	opts module.Options
}

func (e *App) Run() error {
	return nil
}
