package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CheckTokenRepo interface {
	GetToken(ctx context.Context, req interface{}) error
	CheckToken(ctx context.Context, params interface{}, token string) (bool, error)
}

type CheckTokenUsecase struct {
	log  *log.Helper
	repo CheckTokenRepo
}

func NewCheckTokenUsecase(logger log.Logger) *CheckTokenUsecase {
	return &CheckTokenUsecase{
		log: log.NewHelper(logger),
	}
}

func (uc CheckTokenUsecase) CheckToken(ctx context.Context, request interface{}, token string) (bool, error) {
	return uc.repo.CheckToken(ctx, request, token)
}

func (uc CheckTokenUsecase) GetToken(ctx context.Context, request interface{}) error {
	return uc.repo.GetToken(ctx, request)
}
