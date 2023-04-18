// Package cron 定时任务
package cron

import (
	"context"
	"encoding/json"
	"github.com/ykstudy/gocore/utils"
	"io/ioutil"

	"github.com/ykstudy/gocore/utils/logging"
	"github.com/ykstudy/stock/datacenter"
	"github.com/ykstudy/stock/models"
)

// SyncFundManagers 同步基金经理
func SyncFundManagers() {
	if !utils.IsTradingDay() {
		return
	}
	ctx := context.Background()
	managers, err := datacenter.EastMoney.FundMangers(ctx, "all", "penavgrowth", "desc")
	if err != nil {
		logging.Error(ctx, "SyncFundManagers error:"+err.Error())
	}
	managers.SortByYieldse()
	if len(managers) != 0 {
		models.FundManagers = managers
	}

	// 更新文件
	b, err := json.MarshalIndent(managers, "", "  ")
	if err != nil {
		logging.Errorf(ctx, "SyncFundManagers json marshal error:", err)
		promSyncError.WithLabelValues("SyncFundManagers").Inc()
		return
	}
	if err := ioutil.WriteFile(models.FundManagersFilename, b, 0666); err != nil {
		logging.Errorf(ctx, "SyncFundManagers WriteFile error:", err)
		promSyncError.WithLabelValues("SyncFundManagers").Inc()
		return
	}
}
