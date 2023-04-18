//go:generate swag init --dir ./ --generalInfo routes/routes.go --propertyStrategy snakecase --output ./routes/docs

// Package main stock is my stock bot
package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"time"

	"github.com/spf13/viper"
	"github.com/ykstudy/stock/cmds"
	"github.com/ykstudy/stock/models"
	"github.com/ykstudy/stock/version"
)

var (
	// DefaultLoglevel 日志级别默认值
	DefaultLoglevel = "info"
	// ProcessorOptions 要启动运行的进程可选项
	ProcessorOptions = []string{cmds.ProcessorChecker, cmds.ProcessorExportor, cmds.ProcessorWebserver, cmds.ProcessorIndex, cmds.ProcessorJSON}
)

func init() {
	viper.SetDefault("app.chan_size", 1)
	models.InitGlobalVars()
}

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "stock"
	app.Usage = "ctt521 的股票工具程序"
	app.UsageText = `该程序不构成任何投资建议，程序只是个人辅助工具，具体分析仍然需要自己判断。

官网地址: http://stock.ctt521.com`
	app.Version = version.Version
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		{
			Name:  "ctt521",
			Email: "kuwen1249163@163.com",
		},
	}
	app.Copyright = "(c) 2023 ctt521"

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show the version",
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "loglevel",
			Aliases:     []string{"l"},
			Value:       DefaultLoglevel,
			Usage:       "cmd 日志级别 [debug|info|warn|error]",
			EnvVars:     []string{"INVESTOOL_CMD_LOGLEVEL"},
			DefaultText: DefaultLoglevel,
		},
	}
	app.BashComplete = func(c *cli.Context) {
		if c.NArg() > 0 {
			return
		}
		for _, i := range ProcessorOptions {
			fmt.Println(i)
		}
	}

	app.Commands = append(app.Commands, cmds.CommandExportor())
	app.Commands = append(app.Commands, cmds.CommandChecker())
	app.Commands = append(app.Commands, cmds.CommandWebserver())
	app.Commands = append(app.Commands, cmds.CommandIndex())
	app.Commands = append(app.Commands, cmds.CommandJSON())

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
	}

}
