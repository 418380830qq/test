package router

import (
	"github.com/gin-gonic/gin"
	"kainengpaichong/service"
	"kainengpaichong/until"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(until.OptionsHandler) //crog跨域
	r.POST("/selectallfile", service.SelectAllfile)
	r.POST("/fetch", func(c *gin.Context) {
		// 调用 service.Fetch() 函数执行
		service.Fetch()
		c.JSON(http.StatusOK, gin.H{
			"message": "执行成功",
		})
	})
	r.POST("/update", service.UpdataMaxprice)
	r.POST("/updatenowprice", service.UpdataNowPrice)
	r.POST("/updateStatus", service.UpdataStatus)
	r.POST("/updateupbysuggest", service.UpdateUpBySuggest)
	return r
}
