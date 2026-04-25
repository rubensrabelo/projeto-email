package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]Campaign), args.Error(1)
}

func (r *repositoryMock) GetBy(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), args.Error(1)
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi!",
		Emails:  []string{"teste1@test.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(repositoryMock)
	repoMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repoMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.NotNil(err)
	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	repoMock := new(repositoryMock)
	repoMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)
	service.Repository = repoMock

	service.Create(newCampaign)

	repoMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(repositoryMock)
	repoMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repoMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetBy_Campaign(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(repositoryMock)
	campaignExpected, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repoMock.On("GetBy", campaignExpected.ID).Return(campaignExpected, nil)
	service.Repository = repoMock

	campaignResponse, err := service.GetBy(campaignExpected.ID)

	assert.Nil(err)
	assert.Equal(campaignExpected.ID, campaignResponse.ID)
	assert.Equal(campaignExpected.Name, campaignResponse.Name)
}

func Test_GetBy_Error(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(repositoryMock)
	repoMock.On("GetBy", mock.Anything).Return(nil, errors.New("database error"))
	service.Repository = repoMock

	_, err := service.GetBy("invalid_id")

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("Something wrong'"))
	service.Repository = repositoryMock

	_, err := service.GetBy(campaign.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}