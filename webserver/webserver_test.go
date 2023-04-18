package webserver

import (
	"github.com/ykstudy/gocore/utils"
	"testing"

	"github.com/spf13/viper"
)

func TestViperConfig(t *testing.T) {
	InitWithConfigFile("../config.toml")
	defer viper.Reset()
	if !utils.IsInitedViper() {
		t.Error("init viper failed")
	}
}
