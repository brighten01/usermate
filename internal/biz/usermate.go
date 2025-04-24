package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"usermate/internal/data/elasticsearch"
	"usermate/internal/data/elasticsearch/mates"
	"usermate/internal/data/model"
)

type AddMateRequest struct {
	UserName string `json:"username"`
	GroupId  int32  `json:"group_id"`
	RealName string `json:"real_name"`
	Tags     string `json:"tags"`
	Birthday string `json:"birthday"`
	Hobby    string `json:"hobby"`
	Nickname string `json:"nickname"`
	Images   string `json:"images"`
	Age      int32  `json:"age"`
	Province string `json:"province"`
	Sign     string `json:"sign"`
	VideoUrl string `json:"videourl"`
}

type AddMateResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type UpdateMateRequest struct {
	UserName string `json:"username"`
	GroupId  int32  `json:"group_id"`
	RealName string `json:"real_name"`
	Tags     string `json:"tags"`
	Birthday string `json:"birthday"`
	Hobby    string `json:"hobby"`
	Nickname string `json:"nickname"`
	Images   string `json:"images"`
	Age      int32  `json:"age"`
	Province string `json:"province"`
	Sign     string `json:"sign"`
	VideoUrl string `json:"videourl"`
}

type SearchUserMate struct {
	UserName string `json:"username"`
	GroupId  int32  `json:"group_id"`
	RealName string `json:"real_name"`
	Tags     string `json:"tags"`
	Birthday string `json:"birthday"`
	Hobby    string `json:"hobby"`
	Nickname string `json:"nickname"`
	Images   string `json:"images"`
	Age      int32  `json:"age"`
	Province string `json:"province"`
	Sign     string `json:"sign"`
	VideoUrl string `json:"videourl"`
}

type CreateOrderInfo struct {
	Uid             int64     `json:"uid"`
	UserMateId      int64     `json:"user_mate_id"`
	ServiceCategory string    `json:"service_category"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	Amount          float64   `json:"amount"`
	Discount        float64   `json:"discount"`
	Avatar          string    `json:"avatar"`
	LinkUrl         string    `json:"link_url"`
	IsOrderAfter    int8      `json:"is_order_after"`
	Gender          int8      `json:"gender"`
	Level           int8      `json:"level"`
	Duration        int8      `json:"duration"`
	Wechat          string    `json:"wechat"`
	Note            string    `json:"note"`
	Payment         int64     `json:"payment"`
}

type UpdateOrderInfo struct {
	OrderId string `json:"order_id"`
	Status  int8   `json:"status"`
	Note    string `json:"note"`
}

type LevelRequest struct {
	Level     int32  `json:"level"`
	LevelName string `json:"level_name"`
	Status    int32  `json:"status"`
	Radio     int32  `json:"radio"`
}

type ServiceCategoryRequest struct {
	CategoryName string `json:"category_name"`
	BaseAmount   int32  `json:"base_amount"`
	ParentId     int32  `json:"parent_id"`
	Status       int32  `json:"status"`
}
type UserMateRepo interface {
	AddUserMate(ctx context.Context, addmate AddMateRequest) (AddMateResponse, error)
	UserMateList(ctx context.Context, pageno int32, pagesize int32) ([]*model.UserMate, error)
	UserMateDetail(ctx context.Context, id int32) (*model.UserMate, error)
	DeleteUserMate(ctx context.Context, id int32) error
	UpdateUserMate(ctx context.Context, id int32, updateMate *UpdateMateRequest) error
	SearchUserMate(ctx context.Context, username string) ([]*model.UserMate, error)
	CreateOrder(ctx context.Context, orderCreate *CreateOrderInfo) (id int64, order_id string, err error)
	UpdateOrder(ctx context.Context, orderUpdate *UpdateOrderInfo) (id int64, order_id string, err error)
	OrderList(ctx context.Context, pageno int32, pagesize int32) ([]*model.OrderList, error)
	OrderDetail(ctx context.Context, order_id string) (*model.OrderList, error)
	AddLevel(ctx context.Context, levelinfo *LevelRequest) error
	AddServiceCategory(ctx context.Context, category *ServiceCategoryRequest) error
	CreateToKafka(ctx context.Context, orderInfo *model.OrderList) error
	GetOrderInfoById(ctx context.Context, order_id string) (*model.OrderList, error)
}

type UserMateUsecase struct {
	repo     UserMateRepo
	log      *log.Helper
	esclient *elasticsearch.ES
}

func NewUserMateUsecase(repo UserMateRepo, logger log.Logger, es *elasticsearch.ES) *UserMateUsecase {
	return &UserMateUsecase{repo: repo, log: log.NewHelper(logger), esclient: es}
}

func (uc *UserMateUsecase) AddUserMate(ctx context.Context, addmate AddMateRequest) (*AddMateResponse, error) {
	uc.log.WithContext(ctx).Infof("AddUserMate: %v", addmate)
	addUsermateResponse, err := uc.repo.AddUserMate(ctx, addmate)
	if err != nil {
		return nil, err
	}

	addMateDTO := mates.AddMateDTO{
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
		UpdatedAt: time.Now().Format("2006-01-02 03:04:05"),
		CreatedAt: time.Now().Format("2006-01-02 03:04:05"),
	}
	err = uc.esclient.UserMateClient.AddUserMates(ctx, addMateDTO)
	if err != nil {
		//只做记录不影响搜索
		uc.log.WithContext(ctx).Errorf("sync to es data failed %v", addMateDTO)
	}
	return &addUsermateResponse, nil
}

func (uc *UserMateUsecase) UserMateList(ctx context.Context, pageno int32, pagesize int32) ([]*model.UserMate, error) {
	uc.log.WithContext(ctx).Infof("UserMateList: %v", pageno, pagesize)
	return uc.repo.UserMateList(ctx, pageno, pagesize)
}

// 获取用户详情
func (uc *UserMateUsecase) UserMateDetail(ctx context.Context, id int32) (*model.UserMate, error) {
	uc.log.WithContext(ctx).Infof("UserMateDetail: %v", id)
	return uc.repo.UserMateDetail(ctx, id)
}

func (uc *UserMateUsecase) DeleteUserMate(ctx context.Context, id int32) error {
	uc.log.WithContext(ctx).Infof("DeleteUserMate: %v", id)
	return uc.repo.DeleteUserMate(ctx, id)
}

func (uc *UserMateUsecase) UpdateUserMate(ctx context.Context, id int32, updateMate *UpdateMateRequest) error {
	uc.log.WithContext(ctx).Infof("UpdateUserMate: %v", id)
	return uc.repo.UpdateUserMate(ctx, id, updateMate)
}

// 按照用户名搜索用户
func (uc *UserMateUsecase) SearchMates(ctx context.Context, username string) ([]*mates.UserMateVo, error) {
	uc.log.WithContext(ctx).Infof("user searching usermate %v ", username)
	searchDTO := mates.SearchMateDTO{}
	_, err := uc.esclient.UserMateClient.SearchMates(ctx, searchDTO)
	if err != nil {
		return nil, err
	}
	return []*mates.UserMateVo{}, nil
}

func (uc *UserMateUsecase) CreateOrderInfo(ctx context.Context, info *CreateOrderInfo) (id int64, order_id string, err error) {
	uc.log.WithContext(ctx).Infof("Create Order Info %v", info)
	return uc.repo.CreateOrder(ctx, info)
}

func (uc *UserMateUsecase) UpdateOrderInfo(ctx context.Context, info *UpdateOrderInfo) (id int64, order_id string, err error) {
	uc.log.WithContext(ctx).Infof("Update Order Info %v", info)
	return uc.repo.UpdateOrder(ctx, info)
}

func (uc *UserMateUsecase) OrderList(ctx context.Context, pageno int32, pagesize int32) ([]*model.OrderList, error) {
	uc.log.WithContext(ctx).Infof("OrderList: %v", pageno, pagesize)
	return uc.repo.OrderList(ctx, pageno, pagesize)
}

func (uc *UserMateUsecase) OrderDetail(ctx context.Context, order_id string) (*model.OrderList, error) {
	uc.log.WithContext(ctx).Infof("OrderDetail: %v", order_id)
	return uc.repo.OrderDetail(ctx, order_id)
}

func (uc *UserMateUsecase) AddLevel(ctx context.Context, request *LevelRequest) error {
	uc.log.WithContext(ctx).Infof("level: %v", request)
	return uc.repo.AddLevel(ctx, request)
}

func (uc *UserMateUsecase) AddServiceCategory(ctx context.Context, category *ServiceCategoryRequest) error {
	uc.log.WithContext(ctx).Infof("level: %v", category)
	return uc.repo.AddServiceCategory(ctx, category)
}

func (uc *UserMateUsecase) CreateOrderToKafka(ctx context.Context, orderInfo *model.OrderList) error {
	uc.log.WithContext(ctx).Infof("order to es %v ", orderInfo)
	return uc.repo.CreateToKafka(ctx, orderInfo)
}

func (uc *UserMateUsecase) GetOrderInfoById(ctx context.Context, order_id string) (*model.OrderList, error) {
	return uc.repo.GetOrderInfoById(ctx, order_id)
}
