package routes

import (
	"github.com/ykstudy/gocore/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	viper.Set("basic_auth.username", "admin")
	viper.Set("basic_auth.password", "admin")
	defer viper.Reset()
	Register(r)
	recorder, err := utils.RequestHTTPHandler(
		r,
		"GET",
		"/x/ping",
		nil,
		map[string]string{"Authorization": "Basic YWRtaW46YWRtaW4="},
	)
	assert.Nil(t, err)
	assert.Equal(t, recorder.Code, 200)
}
