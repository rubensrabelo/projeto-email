package endpoints

import (
	"emailn/internal/contract"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (r *serviceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (r *serviceMock) Get() (interface{}, error) {
	args := r.Called()
	return args.Get(0), args.Error(1)
}
