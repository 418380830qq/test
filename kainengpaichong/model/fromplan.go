package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Fromplan struct {
	gorm.Model
	Planid          int
	PlanName        string
	Wordid          int `gorm:"primaryKey;unique"`
	Word            string
	MinPrice        float64 // 以字符串形式存储 JSON 数组
	Maxprice        float64
	SuggestPriceStr float64
	NowPrice        float64
	Setprice        float64
	Remarks         string
}

func NewFromplan() *Fromplan {
	return &Fromplan{
		Setprice: 20.0,
	}
}
func (plan *Fromplan) TableName() string {
	return "Fromplan"
}
func CreateOrUpdateData(data *Fromplan) {
	newData := &Fromplan{}
	// 使用 FirstOrCreate 方法查找匹配记录或创建新记录
	if err := DB.Where(Fromplan{Wordid: data.Wordid}).Assign(data).FirstOrCreate(newData).Error; err != nil {
		fmt.Println("创建或更新数据时发生错误:", err)
	}
}

func SelectAllFiles() ([]*Fromplan, error) {
	// 执行数据库查询
	var results []*Fromplan
	if err := DB.Find(&results).Error; err != nil {
		// 处理查询过程中的错误
		return nil, err
	}
	// 返回查询结果0
	return results, nil
}

func UpdateMaxPrice(data *Fromplan) *gorm.DB {
	return DB.Model(data).Updates(Fromplan{Wordid: data.Wordid, Setprice: data.Setprice})
}

func SelectMaxPrice(data *Fromplan) (float64, error) {
	var maxPrice float64
	result := DB.Model(&Fromplan{}).Where("wordid = ?", data.Wordid).Select("setprice").Row()

	if err := result.Scan(&maxPrice); err != nil {
		return 0, err
	}

	return maxPrice, nil
}
