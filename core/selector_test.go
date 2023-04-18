package core

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ykstudy/gocore/utils/logging"
	"github.com/ykstudy/stock/datacenter/eastmoney"
)

func TestAutoFilterStocks(t *testing.T) {
	logging.SetLevel("error")
	checker := NewChecker(_ctx, DefaultCheckerOptions)
	s := NewSelector(_ctx, eastmoney.DefaultFilter, checker)
	_, err := s.AutoFilterStocks(_ctx)
	require.Nil(t, err)
	// t.Log(result)
}
