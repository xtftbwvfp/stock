// Package cron 定时任务
package cron

import (
	"context"
	"github.com/ykstudy/gocore/utils"

	"github.com/ykstudy/stock/datacenter"
	"github.com/ykstudy/stock/models"
)

// SyncBond 同步债券
func SyncBond() {
	if !utils.IsTradingDay() {
		return
	}
	ctx := context.Background()
	syl := datacenter.ChinaBond.QueryAAACompanyBondSyl(ctx)
	if syl != 0 {
		models.AAACompanyBondSyl = syl
	}
}
