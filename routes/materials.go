// 学习资料页面

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/ykstudy/stock/version"
)

// MaterialItem 学习资料具体信息
type MaterialItem struct {
	Name        string `json:"name"`
	DownloadURL string `json:"download_url"`
	Desc        string `json:"desc"`
}

// MaterialSeries 某一个系列的资料
//
//	{
//	    "飙股在线等": [
//	        MaterialItem, ...
//	    ]
//	}
type MaterialSeries map[string][]MaterialItem

// TypedMaterialSeries 对MaterialSeries进行分类，如：视频、电子书等
//
//	{
//	    "videos": [
//	        MaterialSeries, ...
//	    ],
//	    "ebooks": [
//	        MaterialSeries, ...
//	    ]
//	}
type TypedMaterialSeries map[string][]MaterialSeries

// AllMaterialsList 包含全部资料信息的大JSON列表
// [
//
//	TypedMaterialSeries, ...
//
// ]
type AllMaterialsList []TypedMaterialSeries

// Materials godoc
func Materials(c *gin.Context) {
	data := gin.H{
		"Env":       viper.GetString("env"),
		"Version":   version.Version,
		"PageTitle": "Stock | 资料",
	}
	var mlist AllMaterialsList
	data["AllMaterialsList"] = mlist
	c.HTML(http.StatusOK, "materials", data)
	return
}
