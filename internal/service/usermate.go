package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"time"
	pb "usermate/api/usermate/v1"
	"usermate/internal/biz"
)

type UserMateService struct {
	pb.UnimplementedUserMateServer
	uc  *biz.UserMateUsecase
	log *log.Helper
}

func NewUserMateService(uc *biz.UserMateUsecase, logger log.Logger) *UserMateService {
	return &UserMateService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserMateService) AddUserMate(ctx context.Context, req *pb.UserMateRequest) (*pb.UserMateReply, error) {
	usermate := &biz.AddMateRequest{
		UserName: req.Username,
		GroupId:  req.GroupId,
		RealName: req.RealName,
		Tags:     req.Tags,
		Birthday: req.Birthday,
		Hobby:    req.Hobby,
		Nickname: req.Nickname,
		Images:   req.Images,
		Age:      req.Age,
		Province: req.Province,
		Sign:     req.Sign,
		VideoUrl: req.Videourl,
	}

	// 调用业务逻辑处理添加用户
	response, err := s.uc.AddUserMate(ctx, *usermate)
	if err != nil {
		return &pb.UserMateReply{
			Reply: &pb.Reply{
				Code:    500,
				Message: err.Error(),
			},
		}, err
	}

	return &pb.UserMateReply{
		Reply: &pb.Reply{
			Code:    response.Code,
			Message: response.Message,
		},
	}, nil
}

// 删除用户
func (s *UserMateService) DeleteUserMate(ctx context.Context, req *pb.DeleteMateRequest) (*pb.DeleteMateReply, error) {
	err := s.uc.DeleteUserMate(ctx, req.Id)
	if err != nil {
		return &pb.DeleteMateReply{
			Reply: &pb.Reply{
				Code:    206,
				Message: "failed to delete user mate",
			},
		}, err
	}
	return &pb.DeleteMateReply{
		Reply: &pb.Reply{
			Code:    200,
			Message: "success",
		},
	}, nil
}

func (s *UserMateService) UserMateDetail(ctx context.Context, req *pb.UserMateShowRequest) (*pb.UserMateShowReply, error) {
	response, err := s.uc.UserMateDetail(ctx, req.Id)
	if err != nil {
		return &pb.UserMateShowReply{}, err
	}

	return &pb.UserMateShowReply{
		Mateinfo: &pb.UserMateInfo{
			Username:  response.UserName,
			GroupId:   response.GroupId,
			RealName:  response.RealName,
			Tags:      response.Tags,
			Birthday:  response.Birthday,
			Hobby:     response.Hobby,
			Nickname:  response.Nickname,
			Images:    response.Images,
			Age:       response.Age,
			Province:  response.Province,
			Sign:      response.Sign,
			Videourl:  response.VideoUrl,
			CreatedAt: response.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:  response.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

func (s *UserMateService) UpdateUserMate(ctx context.Context, req *pb.UserMateUpdateRequest) (*pb.UserMateUpdateReply, error) {
	updateMate := &biz.UpdateMateRequest{
		UserName: req.Username,
		GroupId:  req.GroupId,
		RealName: req.RealName,
		Tags:     req.Tags,
		Birthday: req.Birthday,
		Hobby:    req.Hobby,
		Nickname: req.Nickname,
		Images:   req.Images,
		Age:      req.Age,
		Province: req.Province,
		Sign:     req.Sign,
		VideoUrl: req.Videourl,
	}
	err := s.uc.UpdateUserMate(ctx, req.Uid, updateMate)
	if err != nil {
		return &pb.UserMateUpdateReply{

			Code:    206,
			Message: "failed to update user mate",
		}, err
	}
	return &pb.UserMateUpdateReply{
		Code:    200,
		Message: "success",
	}, nil
}

func (s *UserMateService) ListUserMate(ctx context.Context, req *pb.ListMateRequest) (*pb.ListMateResponse, error) {
	response, err := s.uc.UserMateList(ctx, req.Page, req.Pagesize)
	s.log.WithContext(ctx).Infof("page is %v , pagesize is %v", req.Page, req.Pagesize)
	if err != nil {
		return &pb.ListMateResponse{}, err
	}
	userMateList := make([]*pb.UserMateInfo, 0)
	for _, userMate := range response {
		userMateList = append(userMateList, &pb.UserMateInfo{
			Username:  userMate.UserName,
			GroupId:   userMate.GroupId,
			RealName:  userMate.RealName,
			Tags:      userMate.Tags,
			Birthday:  userMate.Birthday,
			Hobby:     userMate.Hobby,
			Nickname:  userMate.Nickname,
			Images:    userMate.Images,
			Age:       userMate.Age,
			Province:  userMate.Province,
			Sign:      userMate.Sign,
			Videourl:  userMate.VideoUrl,
			UpdateAt:  userMate.UpdatedAt.Format("2006-01-02 15:04:05"),
			CreatedAt: userMate.CreateAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.ListMateResponse{
		List: userMateList,
	}, nil
}

func (s *UserMateService) SearchUserMate(ctx context.Context, req *pb.SearchUserMateRequest) (*pb.SearchUserMateResponse, error) {
	userMate, err := s.uc.SearchMates(ctx, req.Name)

	if err != nil {
		s.log.WithContext(ctx).Infof("search err  is %v", err)
		return &pb.SearchUserMateResponse{}, nil
	}
	var mateList = make([]*pb.UserMateInfo, 0)
	for _, usermate := range userMate {
		mateList = append(mateList, &pb.UserMateInfo{
			Username:  usermate.UserName,
			GroupId:   usermate.GroupId,
			Tags:      usermate.Tags,
			Birthday:  usermate.Birthday,
			Hobby:     usermate.Hobby,
			Nickname:  usermate.Nickname,
			Images:    usermate.Images,
			Age:       usermate.Age,
			Province:  usermate.Province,
			Sign:      usermate.Sign,
			Videourl:  usermate.VideoUrl,
			UpdateAt:  usermate.UpdatedAt.Format("2006-01-02 15:04:05"),
			CreatedAt: usermate.CreateAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &pb.SearchUserMateResponse{
		Mateinfo: mateList,
	}, nil
}

func (s *UserMateService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	CreateOrderInfo := &biz.CreateOrderInfo{
		Uid:             int64(req.Uid),
		UserMateId:      int64(req.UserMateId),
		ServiceCategory: strconv.Itoa(int(req.ServiceCategory)),
		StartTime:       func() time.Time { t, _ := time.Parse("2006-01-02 15:04:05", req.StartTime); return t }(),
		EndTime:         func() time.Time { t, _ := time.Parse("2006-01-02 15:04:05", req.EndTime); return t }(),
		Amount:          float64(req.Amount),
		Discount:        float64(req.Discount),
		Avatar:          req.Avatar,
		LinkUrl:         req.LinkUrl,
		IsOrderAfter:    int8(req.IsOrderAfter),
		Gender:          int8(req.Gender),
		Level:           int8(req.Level),
		Duration:        int8(req.Duration),
		Wechat:          req.Wechat,
		Note:            req.Note,
		Payment:         int64(req.Payment),
	}
	_, _, err := s.uc.CreateOrderInfo(ctx, CreateOrderInfo)
	if err != nil {
		return &pb.CreateOrderReply{}, nil
	}
	return &pb.CreateOrderReply{
		Code:         200,
		Message:      "success",
		UserMateId:   int32(CreateOrderInfo.UserMateId),
		UserMateName: "",
	}, nil
}

func (s *UserMateService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderReply, error) {
	UpdateOrderInfo := &biz.UpdateOrderInfo{
		OrderId: req.OrderId,
		Status:  int8(req.Status),
	}
	_, _, err := s.uc.UpdateOrderInfo(ctx, UpdateOrderInfo)
	if err != nil {
		return &pb.UpdateOrderReply{}, nil
	}
	return &pb.UpdateOrderReply{
		Code:    200,
		Message: "success",
	}, nil
}

func (s *UserMateService) OrderList(ctx context.Context, req *pb.OrderListRequest) (*pb.OrderListResponse, error) {
	response, err := s.uc.OrderList(ctx, req.CustomerId, 10)
	if err != nil {
		return &pb.OrderListResponse{}, err
	}
	orderList := make([]*pb.OrderDetailResponse, 0)
	for _, order := range response {
		orderList = append(orderList, &pb.OrderDetailResponse{
			OrderId:             order.OrderID,
			Uid:                 int32(order.UID),
			UserMateId:          int32(order.UserMateID),
			ServiceCategory:     func() int32 { v, _ := strconv.Atoi(order.ServiceCategory); return int32(v) }(),
			StartTime:           order.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:             order.EndTime.Format("2006-01-02 15:04:05"),
			Amount:              order.Amount,
			Payment:             int32(order.Payment),
			Avatar:              order.Avatar,
			LinkUrl:             order.LinkURL,
			IsOrderAfter:        int32(order.IsOrderAfter),
			Gender:              int32(order.Gender),
			Level:               int32(order.Level),
			Duration:            int32(order.Duration),
			ServiceCategoryName: order.ServiceCategory,
			ServiceCategoryId:   int32(order.ServiceCategoryID),
			Wechat:              order.Wechat,
			Note:                order.Note,
		})
	}
	return &pb.OrderListResponse{
		Data: orderList,
	}, nil
}

func (s *UserMateService) OrderDetail(ctx context.Context, req *pb.OrderDetailRequest) (*pb.OrderDetailResponse, error) {
	response, err := s.uc.OrderDetail(ctx, req.OrderId)
	if err != nil {
		return &pb.OrderDetailResponse{}, err
	}
	return &pb.OrderDetailResponse{
		OrderId:             response.OrderID,
		ServiceCategory:     int32(response.ServiceCategoryID),
		StartTime:           response.CreatedAt.Format("2006-01-02 15:04:05"),
		EndTime:             response.UpdatedAt.Format("2006-01-02 15:04:05"),
		Amount:              float64(response.ServiceCategoryID),
		ServiceCategoryName: response.ServiceCategoryName,
		Gender:              int32(response.Gender),
		Level:               int32(response.Level),
		Duration:            int32(response.Duration),
		Note:                response.Note,
		Payment:             int32(response.Payment),
		Uid:                 int32(response.UID),
		UserMateId:          int32(response.UserMateID),
		Nickname:            response.Nickname,
		Avatar:              response.Avatar,
		LinkUrl:             response.LinkURL,
	}, nil
}

func (s *UserMateService) AddLevel(ctx context.Context, level *pb.LevelRequest) (*pb.LevelResponse, error) {
	levelRequest := &biz.LevelRequest{
		Level:     level.Level,
		LevelName: level.LevelName,
		Status:    level.Status,
		Radio:     level.Radios,
	}
	error := s.uc.AddLevel(ctx, levelRequest)
	if error != nil {
		return &pb.LevelResponse{}, nil
	}
	return &pb.LevelResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (s *UserMateService) AddServiceCategory(ctx context.Context, category *pb.ServiceCategoryRequest) (*pb.ServiceCategoryResponse, error) {
	categoryRequest := &biz.ServiceCategoryRequest{
		CategoryName: category.CategoryName,
		BaseAmount:   category.BaseAmount,
		ParentId:     category.ParentId,
		Status:       category.Status,
	}
	err := s.uc.AddServiceCategory(ctx, categoryRequest)
	if err != nil {
		return &pb.ServiceCategoryResponse{
			Code:    500,
			Message: "failed",
		}, err
	}
	return &pb.ServiceCategoryResponse{
		Code:    200,
		Message: "success",
	}, nil
}
