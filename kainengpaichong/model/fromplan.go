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
	UpBySuggest     float64
	Setprice        float64
	Remarks         string
	Status          bool
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
func UpdateNowPrice(data *Fromplan) *gorm.DB {
	fmt.Println(data.NowPrice)
	return DB.Model(data).Updates(Fromplan{Wordid: data.Wordid, NowPrice: data.NowPrice})
}

func SelectMaxPrice(data *Fromplan) (float64, error) {
	var maxPrice float64
	result := DB.Model(&Fromplan{}).Where("wordid = ?", data.Wordid).Select("setprice").Row()

	if err := result.Scan(&maxPrice); err != nil {
		return 0, err
	}

	return maxPrice, nil
}
func UpdateStatus(wordid int, newStatus bool) error {
	// 执行更新操作，这里假设您使用了 GORM 或类似的 ORM 库
	var word Fromplan
	result := DB.Where("wordid = ?", wordid).First(&word)
	if result.Error != nil {
		// 处理查询失败的情况，可以打印错误信息或者进行其他处理
		return result.Error
	}

	// 更新状态字段为 newStatus
	word.Status = newStatus
	result = DB.Save(&word)
	if result.Error != nil {
		// 处理更新失败的情况，可以打印错误信息或者进行其他处理
		return result.Error
	}

	return nil
}

func SelectStatus(wordid int) (bool, error) {
	var word Fromplan
	result := DB.Where("wordid = ?", wordid).First(&word)
	if result.Error != nil {
		// 处理查询失败的情况，可以打印错误信息或者进行其他处理
		fmt.Println("查询数据库失败:", result.Error)
		return false, result.Error
	}

	// 返回词的状态
	return word.Status, nil
}

func SelectByWordid(wordid int) (Fromplan, error) {
	var word Fromplan
	result := DB.Where("wordid = ?", wordid).First(&word)
	if result.Error != nil {
		// 处理查询失败的情况，例如打印日志或返回特定错误信息
		return Fromplan{}, result.Error
	}

	return word, nil
}

func UpdateUpBysuggest(wordid int, UpBysuggest float64) error {
	// 执行更新操作，这里假设您使用了 GORM 或类似的 ORM 库
	var word Fromplan
	result := DB.Where("wordid = ?", wordid).First(&word)
	if result.Error != nil {
		// 处理查询失败的情况，可以打印错误信息或者进行其他处理
		return result.Error
	}

	// 更新状态字段为 newStatus
	word.UpBySuggest = UpBysuggest
	result = DB.Save(&word)
	if result.Error != nil {
		// 处理更新失败的情况，可以打印错误信息或者进行其他处理
		return result.Error
	}

	return nil
}
