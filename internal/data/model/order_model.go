package model

import (
	"time"
)

type Order struct {
	ID              int64     `gorm:"column:id;comment:主键ID"`
	OrderID         string    `gorm:"column:order_id;size:100;comment:订单id"`
	UID             int64     `gorm:"column:uid;comment:用户id"`
	UserMateID      int64     `gorm:"column:user_mate_id;comment:接单人id"`
	ServiceCategory string    `gorm:"column:service_category;comment:订单类型 各种服务类型"`
	StartTime       time.Time `gorm:"column:start_time;comment:开始时间"`
	EndTime         time.Time `gorm:"column:end_time;comment:结束时间"`
	Status          int8      `gorm:"column:status;comment:订单状态 1开始 2进行中 3已完成 4退单 5取消 6关闭"`
	Amount          float64   `gorm:"column:amount;type:decimal(10,2);comment:订单金额"`
	CreatedAt       time.Time `gorm:"column:createdAt;comment:下单时间"`
	Payment         int64     `gorm:"column:payment;comment:1余额支付 2支付宝 3微信"`
	Discount        float64   `gorm:"column:discount;type:decimal(10,2);comment:优惠券扣减"`
	UpdatedAt       time.Time `gorm:"column:updatedAt;comment:更新时间"`
	Avatar          string    `gorm:"column:avatar;size:100;comment:图片地址"`
	LinkURL         string    `gorm:"column:link_url;size:100;comment:主页地址"`
	IsOrderAfter    int8      `gorm:"column:is_order_after;comment:是否续单 1续 2不续"`
}

type OrderDetail struct {
	ID                  int       `gorm:"primaryKey;column:id;autoIncrement;comment:主键ID"`
	OrderID             string    `gorm:"column:order_id;size:100;comment:主订单id"`
	Gender              int8      `gorm:"column:gender;comment:1男 2女"`
	Level               int8      `gorm:"column:level;comment:等级"`
	Duration            int8      `gorm:"column:duration;comment:时长"`
	ServiceCategoryID   int       `gorm:"column:service_category_id;comment:服务id"`
	ServiceCategoryName string    `gorm:"column:service_category_name;comment:服务名称"`
	Wechat              string    `gorm:"column:wechat;size:100;comment:微信等联系方式"`
	Note                string    `gorm:"column:note;size:100;comment:备注"`
	CreatedAt           time.Time `gorm:"column:createdAt;comment:创建时间"`
	UpdatedAt           time.Time `gorm:"column:updatedAt;comment:更新时间"`
}

type OrderList struct {
	ID                  int64     `gorm:"column:id;comment:主键ID"`
	OrderID             string    `gorm:"column:order_id;size:100;comment:订单id"`
	UID                 int64     `gorm:"column:uid;comment:用户id"`
	UserMateID          int64     `gorm:"column:user_mate_id;comment:接单人id"`
	ServiceCategory     string    `gorm:"column:service_category;comment:订单类型 各种服务类型"`
	StartTime           time.Time `gorm:"column:start_time;comment:开始时间"`
	EndTime             time.Time `gorm:"column:end_time;comment:结束时间"`
	Status              int8      `gorm:"column:status;comment:订单状态 1开始 2进行中 3已完成 4退单 5取消 6关闭"`
	Amount              float64   `gorm:"column:amount;type:decimal(10,2);comment:订单金额"`
	Discount            float64   `gorm:"column:discount;type:decimal(10,2);comment:优惠券扣减"`
	Avatar              string    `gorm:"column:avatar;size:100;comment:图片地址"`
	LinkURL             string    `gorm:"column:link_url;size:100;comment:主页地址"`
	IsOrderAfter        int8      `gorm:"column:is_order_after;comment:是否续单 1续 2不续"`
	Gender              int8      `gorm:"column:gender;comment:1男 2女"`
	Level               int8      `gorm:"column:level;comment:等级"`
	Duration            int8      `gorm:"column:duration;comment:时长"`
	ServiceCategoryID   int       `gorm:"column:service_category_id;comment:服务id"`
	ServiceCategoryName string    `gorm:"column:service_category_name;comment:服务名称"`
	Wechat              string    `gorm:"column:wechat;size:100;comment:微信等联系方式"`
	Note                string    `gorm:"column:note;size:100;comment:备注"`
	CreatedAt           time.Time `gorm:"column:createdAt;comment:创建时间"`
	UpdatedAt           time.Time `gorm:"column:updatedAt;comment:更新时间"`
	Payment             int64     `gorm:"column:payment;comment:1余额支付 2支付宝 3微信"`
	Nickname            string
}
