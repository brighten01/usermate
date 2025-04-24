package mates

import "time"

type UserMateVo struct {
	Id        int64     `gorm:"column:id;primary_key;auto_increment"`
	UserName  string    `gorm:"column:username;not null"`
	GroupId   int32     `gorm:"column:group_id;not null"`
	RealName  string    `gorm:"column:real_name;not null"`
	Tags      string    `gorm:"column:tags;not null"`
	Birthday  string    `gorm:"column:birthday;not null"`
	Hobby     string    `gorm:"column:hobby;not null"`
	Nickname  string    `gorm:"column:nickname;not null"`
	Images    string    `gorm:"column:images;not null"`
	Age       int32     `gorm:"column:age;not null"`
	Province  string    `gorm:"column:province;not null"`
	Sign      string    `gorm:"column:sign;not null"`
	VideoUrl  string    `gorm:"column:videourl;not null"`
	CreateAt  time.Time `gorm:"column:createdAt;null"`
	UpdatedAt time.Time `gorm:"column:updatedAt;null"`
}
