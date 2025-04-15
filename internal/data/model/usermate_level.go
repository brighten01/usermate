package model

import "time"

type UserMateLevel struct {
	ID        int       `gorm:"primaryKey;autoIncrement;column:id;comment:主键ID"`
	Level     int8      `gorm:"column:level;comment:等级"`
	LevelName string    `gorm:"type:varchar(100);column:level_name;comment:等级名称"`
	Status    int8      `gorm:"column:status;comment:状态 1 上线 2 下线"`
	Radio     int8      `gorm:"column:radio;comment:系数"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:createdAt;comment:创建时间"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updatedAt;comment:更新时间"`
}
