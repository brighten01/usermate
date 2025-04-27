package data

import (
	"context"
	"usermate/internal/data/model"
)

func (d *Data) TestData(ctx context.Context, nickName string) []*model.UserMate {
	var userlist []*model.UserMate
	result := d.db.WithContext(ctx).Model(&model.UserMate{}).Debug().Where("nick_name =?", nickName).Find(userlist)
	if result.Error != nil {
		return []*model.UserMate{}
	}
	return userlist
}
