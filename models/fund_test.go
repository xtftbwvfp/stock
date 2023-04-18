package models

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ykstudy/stock/datacenter/eastmoney"
)

func TestNewFund(t *testing.T) {
	ctx := context.TODO()
	efund, err := eastmoney.NewEastMoney().QueryFundInfo(ctx, "260104")
	require.Nil(t, err)
	fund := NewFund(ctx, efund)
	b, err := json.Marshal(fund)
	require.Nil(t, err)
	t.Log(string(b))
}
