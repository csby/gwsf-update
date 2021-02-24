package main

import (
	"fmt"
	"github.com/csby/gwsf/gtype"
	"github.com/kardianos/service"
	"os"
	"time"
)

const (
	moduleType    = gtype.SvcUpdModuleType
	moduleName    = gtype.SvcUpdModuleName
	moduleRemark  = "服务更新"
	moduleVersion = "1.0.1.0"
)

var (
	host     = &Host{program: &Program{server: &Server{}}}
	bootTime = gtype.DateTime(time.Now())
)

func init() {
	moduleArgs := &gtype.Args{}
	serverArgs := &gtype.SvcArgs{}
	moduleArgs.Parse(os.Args, moduleType, moduleName, moduleVersion, moduleRemark, serverArgs)
	if serverArgs.Help {
		serverArgs.ShowHelp("", "")
		os.Exit(0)
	}

	if serverArgs.Pkg {
		pkg := &Pkg{binPath: moduleArgs.ModulePath()}
		pkg.Run()
		os.Exit(0)
	}

	cfg := &service.Config{
		Name:        moduleName,
		DisplayName: moduleName,
		Description: moduleRemark,
	}
	svc, err := service.New(host.program, cfg)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(4)
	}
	host.service = svc
	host.program.server.path = moduleArgs.ModulePath()
	serverArgs.Execute(host)
}
