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

// SyncIndustryList 同步行业列表
func SyncIndustryList() {
	if !utils.IsTradingDay() {
		return
	}
	ctx := context.Background()
	indlist, err := datacenter.EastMoney.QueryIndustryList(ctx)
	if err != nil {
		logging.Errorf(ctx, "SyncIndustryList QueryIndustryList error:", err)
		promSyncError.WithLabelValues("SyncIndustryList").Inc()
		return
	}
	if len(indlist) != 0 {
		models.StockIndustryList = indlist
	}

	// 更新文件
	b, err := json.Marshal(indlist)
	if err != nil {
		logging.Errorf(ctx, "SyncIndustryList json marshal error:", err)
		promSyncError.WithLabelValues("SyncIndustryList").Inc()
		return
	}
	if err := ioutil.WriteFile(models.IndustryListFilename, b, 0666); err != nil {
		logging.Errorf(ctx, "SyncIndustryList WriteFile error:", err)
		promSyncError.WithLabelValues("SyncIndustryList").Inc()
		return
	}
}
