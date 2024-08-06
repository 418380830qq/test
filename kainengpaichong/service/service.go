package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kainengpaichong/model"
	"net/http"
	"strconv"
)

type UpdateMaxRequest struct {
	Wordid   int    `json:"wordid"`
	Setprice string `json:"setprice"`
}

type UpdateNowRequest struct {
	Wordid   int    `json:"wordid"`
	Nowprice string `json:"nowprice"`
}
type UpdateStatusRequest struct {
	Wordid int  `json:"wordid"`
	Status bool `json:"status"`
}
type UpdateBysuggestRequest struct {
	Wordid      int    `json:"wordid"`
	UpBysuggest string `json:upbysuggest`
}

func SelectAllfile(c *gin.Context) {
	a, _ := model.SelectAllFiles()
	c.JSON(http.StatusOK, gin.H{
		"message": a,
	})
}

func UpdataMaxprice(c *gin.Context) {
	var updateReq UpdateMaxRequest
	if err := c.BindJSON(&updateReq); err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode request body"})
		return
	}
	setPrice, _ := strconv.ParseFloat(updateReq.Setprice, 64)
	data := model.Fromplan{Wordid: updateReq.Wordid, Setprice: setPrice}
	model.UpdateMaxPrice(&data)
	// 返回响应
	c.JSON(200, gin.H{"message": "UpdateMaxprice request received successfully"})
}

func UpdataNowPrice(c *gin.Context) {
	var word model.Fromplan
	var updateReq UpdateNowRequest
	if err := c.BindJSON(&updateReq); err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode request body"})
		return
	}
	nowprice, _ := strconv.ParseFloat(updateReq.Nowprice, 64)
	data := model.Fromplan{Wordid: updateReq.Wordid, NowPrice: nowprice}
	model.UpdateNowPrice(&data)
	word, _ = model.SelectByWordid(updateReq.Wordid)
	Fetch3(word.Planid, updateReq.Wordid, word.SuggestPriceStr, word.PlanName, word.Maxprice, word.MinPrice, word.Word)
	c.JSON(200, gin.H{"message": "UpdateNowprice request received successfully"})
}

func UpdataStatus(c *gin.Context) {
	var updateReq UpdateStatusRequest
	if err := c.BindJSON(&updateReq); err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode request body"})
		return
	}
	fmt.Println(updateReq)

	model.UpdateStatus(updateReq.Wordid, updateReq.Status)

	c.JSON(200, gin.H{"message": "UpdateStatus request received successfully"})
}
func UpdateUpBySuggest(c *gin.Context) {
	var updateReq UpdateBysuggestRequest
	if err := c.BindJSON(&updateReq); err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode request body"})
		return
	}
	UpBysuggest, _ := strconv.ParseFloat(updateReq.UpBysuggest, 64)
	model.UpdateUpBysuggest(updateReq.Wordid, UpBysuggest)

	c.JSON(200, gin.H{"message": "UpdateStatus request received successfully"})
}
