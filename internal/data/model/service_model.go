package model

import "time"

type ServerCategory struct {
	ID             int       `gorm:"column:id;comment:主键ID"`
	ParentID       int       `gorm:"column:parent_id;comment:上级服务"`
	CategoryName   string    `gorm:"type:varchar(100);column:category_name;comment:分类名称"`
	BaseAmount     string    `gorm:"type:varchar(255);column:base_amount;comment:基础金额"`
	Status         int8      `gorm:"column:status;comment:1开启 2关闭"`
	SevenDaysPrice float64   `gorm:"type:decimal(10,2);column:seven_days_price;comment:7天价格"`
	OneDayPrice    float64   `gorm:"type:decimal(10,2);column:one_day_price;comment:1天价格"`
	MonthPrice     float64   `gorm:"type:decimal(10,2);column:month_price;comment:单月价格"`
	CreatedAt      time.Time `gorm:"autoCreateTime;column:createdAt"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime;column:updatedAt"`
}
