package service

import (
	"github.com/gin-gonic/gin"
	"kainengpaichong/model"
	"net/http"
	"strconv"
)

type UpdateRequest struct {
	Wordid   int    `json:"wordid"`
	Setprice string `json:"setprice"`
}

func SelectAllfile(c *gin.Context) {
	a, _ := model.SelectAllFiles()
	c.JSON(http.StatusOK, gin.H{
		"message": a,
	})
}

func UpdataMaxprice(c *gin.Context) {
	var updateReq UpdateRequest
	if err := c.BindJSON(&updateReq); err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode request body"})
		return
	}
	setPrice, _ := strconv.ParseFloat(updateReq.Setprice, 64)
	data := model.Fromplan{Wordid: updateReq.Wordid, Setprice: setPrice}
	model.UpdateMaxPrice(&data)
	// 返回响应
	c.JSON(200, gin.H{"message": "Update request received successfully"})
}
