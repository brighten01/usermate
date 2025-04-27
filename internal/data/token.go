package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Payload struct {
	alg string `json:"alg"`
	typ string `json:"typ"`
}

type CheckTokenRepo struct {
	log *log.Helper
}

func (c *CheckTokenRepo) CheckToken(ctx context.Context, req interface{}) error {
	c.log.WithContext(ctx).Infof("check user token  req is %v ", req)
	return nil
}

func (c *CheckTokenRepo) GetToken(ctx context.Context, req interface{}) (bool, error) {
	c.log.WithContext(ctx).Infof("get token req is  %v ", req)
	return true, nil
}

func NewCheckTokenRepo(logger log.Logger) *CheckTokenRepo {
	return &CheckTokenRepo{
		log: log.NewHelper(logger),
	}
}
