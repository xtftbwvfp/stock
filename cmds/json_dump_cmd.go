package cmds

import (
	"github.com/urfave/cli/v2"
	"github.com/ykstudy/gocore/utils/logging"
	"github.com/ykstudy/stock/cron"
)

const (
	// ProcessorJSON 导出json数据文件
	ProcessorJSON = "json"
)

// FlagsJSON cli flags
func FlagsJSON() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "dump",
			Aliases: []string{"d"},
			Usage:   "导出json数据文件",
		},
	}
}

// ActionJSON dump json files
func ActionJSON() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		loglevel := c.String("loglevel")
		logging.SetLevel(loglevel)

		if c.Bool("d") {
			cron.SyncFund()
			cron.SyncFundManagers()
			cron.SyncIndustryList()
			return nil
		}
		return nil
	}
}

// CommandJSON dump json files cmd
func CommandJSON() *cli.Command {
	flags := FlagsJSON()
	cmd := &cli.Command{
		Name:   ProcessorJSON,
		Usage:  "JSON数据",
		Flags:  flags,
		Action: ActionJSON(),
	}
	return cmd
}
