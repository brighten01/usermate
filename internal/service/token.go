package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"usermate/internal/biz"
)

type CheckTokenService struct {
	log *log.Helper
	uc  *biz.CheckTokenUsecase
}

func (s *CheckTokenService) CheckToken(ctx context.Context, req interface{}, token string) (bool, error) {
	return s.uc.CheckToken(ctx, req, token)
}

func NewCheckTokenService(logger log.Logger, uc *biz.CheckTokenUsecase) *CheckTokenService {
	return &CheckTokenService{
		log: log.NewHelper(logger),
		uc:  uc,
	}
}
