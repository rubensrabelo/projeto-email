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

func (r *serviceMock) GetBy(id string) (*contract.CampaignResponse, error) {
	args := r.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*contract.CampaignResponse), args.Error(1)
}

