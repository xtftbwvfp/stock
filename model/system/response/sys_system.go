package response

import "github.com/ykstudy/stock/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
