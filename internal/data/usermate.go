package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"time"
	"usermate/internal/biz"
	"usermate/internal/data/model"
	"usermate/pkg/utils"
)

type UserMateRepo struct {
	data *Data
	log  log.Helper
}

func NewUserMateRepo(data *Data, logger log.Logger) biz.UserMateRepo {
	return &UserMateRepo{
		data: data,
		log:  *log.NewHelper(logger),
	}
}

func (u UserMateRepo) AddUserMate(ctx context.Context, addmate biz.AddMateRequest) (biz.AddMateResponse, error) {
	u.log.WithContext(ctx).Infof("AddUserMate: %v", addmate)
	//todo 添加用户请求
	result := u.data.db.Create(&model.UserMate{
		UserName:  addmate.UserName,
		GroupId:   addmate.GroupId,
		RealName:  addmate.RealName,
		Tags:      addmate.Tags,
		Birthday:  addmate.Birthday,
		Hobby:     addmate.Hobby,
		Nickname:  addmate.Nickname,
		Images:    addmate.Images,
		Age:       addmate.Age,
		Province:  addmate.Province,
		Sign:      addmate.Sign,
		VideoUrl:  addmate.VideoUrl,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now().UTC(),
	})

	if result.Error != nil {
		return biz.AddMateResponse{}, result.Error
	}
	return biz.AddMateResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (u UserMateRepo) UserMateList(ctx context.Context, pageno int32, pagesize int32) ([]*model.UserMate, error) {
	u.log.WithContext(ctx).Infof("UserMateList")
	var userMates []*model.UserMate
	result := u.data.db.Limit(int(pagesize)).Offset(int((pageno - 1) * pagesize)).Find(&userMates)
	if result.Error != nil {
		return nil, result.Error
	}
	return userMates, nil
}

// 获取用户详情
func (u UserMateRepo) UserMateDetail(ctx context.Context, id int32) (*model.UserMate, error) {
	u.log.WithContext(ctx).Infof("UserMateDetail: %v", id)
	var userMate model.UserMate
	result := u.data.db.Where("id = ?", id).First(&userMate)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userMate, nil
}

// 删除用户
func (u UserMateRepo) DeleteUserMate(ctx context.Context, id int32) error {
	u.log.WithContext(ctx).Infof("DeleteUserMate: %v", id)
	result := u.data.db.Where("id = ?", id).Delete(&model.UserMate{})
	return result.Error
}

// 更新用户
func (u UserMateRepo) UpdateUserMate(ctx context.Context, id int32, updateMate *biz.UpdateMateRequest) error {
	u.log.WithContext(ctx).Infof("UpdateUserMate: %v", id)
	result := u.data.db.Model(&model.UserMate{}).Where("id = ?", id).Updates(updateMate)
	return result.Error
}

func (u UserMateRepo) SearchUserMate(ctx context.Context, username string) ([]*model.UserMate, error) {
	u.log.WithContext(ctx).Infof("search username %v", username)
	var userMate []*model.UserMate
	result := u.data.db.Where("nickname=?", username).Find(userMate)
	if result.Error != nil {
		return []*model.UserMate{}, result.Error
	}
	return userMate, nil
}

func (u UserMateRepo) CreateOrder(ctx context.Context, createOrderInfo *biz.CreateOrderInfo) (id int64, order_id string, err error) {
	u.log.WithContext(ctx).Infof("user order  %v", createOrderInfo)

	order_id_str := utils.GenerateId()
	//主订单
	var mainOrder *model.Order
	mainOrder = &model.Order{
		OrderID:         order_id,
		UID:             createOrderInfo.Uid,
		UserMateID:      createOrderInfo.UserMateId,
		ServiceCategory: createOrderInfo.ServiceCategory,
		StartTime:       createOrderInfo.StartTime,
		EndTime:         createOrderInfo.EndTime,
		Amount:          createOrderInfo.Amount,
		Payment:         createOrderInfo.Payment,
		Discount:        createOrderInfo.Discount,
		LinkURL:         createOrderInfo.LinkUrl,
		IsOrderAfter:    createOrderInfo.IsOrderAfter,
		UpdatedAt:       time.Now(),
		CreatedAt:       time.Now(),
	}
	result := u.data.db.Create(mainOrder)
	if result.Error != nil {
		return 0, "", err
	}

	//订单详情
	//todo tx 主订单创建失败回滚事务
	//考虑并发因素采取分布式事务方法
	var orderDetail *model.OrderDetail
	orderDetail = &model.OrderDetail{
		OrderID:             order_id,
		Gender:              createOrderInfo.Gender,
		Level:               createOrderInfo.Level,
		Duration:            createOrderInfo.Duration,
		ServiceCategoryID:   1,
		ServiceCategoryName: createOrderInfo.ServiceCategory,
		Wechat:              createOrderInfo.Wechat,
		Note:                createOrderInfo.Note,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
	result = u.data.db.Create(orderDetail)
	if result.Error != nil {
		return 0, "", err
	}
	return mainOrder.ID, order_id_str, nil
}

func (u UserMateRepo) UpdateOrder(ctx context.Context, orderUpdate *biz.UpdateOrderInfo) (id int64, order_id string, err error) {
	u.log.WithContext(ctx).Infof("UpdateOrder: %v", orderUpdate)
	result := u.data.db.Model(&model.Order{}).Where("id = ?", orderUpdate.OrderId).Updates(orderUpdate)
	if result.Error != nil {
		return 0, "", result.Error
	}
	var orderId int64
	orderId, _ = strconv.ParseInt(orderUpdate.OrderId, 10, 64)
	return orderId, orderUpdate.OrderId, nil
}

func (u UserMateRepo) OrderList(ctx context.Context, pageno int32, pagesize int32) ([]*model.OrderList, error) {
	u.log.WithContext(ctx).Infof("OrderList: %v", pageno, pagesize)
	var orders []*model.OrderList

	result := u.data.db.Limit(int(pagesize)).Offset(int((pageno - 1) * pagesize)).Find(&orders)
	return orders, result.Error
}

func (u UserMateRepo) OrderDetail(ctx context.Context, order_id string) (*model.OrderList, error) {
	u.log.WithContext(ctx).Infof("OrderDetail: %v", order_id)
	var orderDetail model.OrderList
	result := u.data.db.Where("order_id = ?", order_id).First(&orderDetail)
	return &orderDetail, result.Error
}
