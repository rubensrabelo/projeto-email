package mock

import (
	"emailn/internal/contract"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (r *ServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (r *ServiceMock) Get() (interface{}, error) {
	args := r.Called()
	return args.Get(0), args.Error(1)
}

func (r *ServiceMock) GetBy(id string) (*contract.CampaignResponse, error) {
	args := r.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*contract.CampaignResponse), args.Error(1)
}

