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

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi",
		Emails:  []string{"teste1@test.com"},
	}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repo := new(repositoryMock)
	repo.On("Save", mock.Anything).Return(nil)
	service := Service{Repository: repo}

	id, err := service.Create(newCampaign)

	assert.Nil(err)
	assert.NotEmpty(id)

	repo.AssertExpectations(t)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	campaignInvalid := newCampaign
	campaignInvalid.Name = ""

	repo := new(repositoryMock)
	service := Service{Repository: repo}

	_, err := service.Create(campaignInvalid)

	assert.NotNil(err)
	assert.Contains(err.Error(), "name") 
	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	assert := assert.New(t)

	repo := new(repositoryMock)
	repo.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return campaign.Name == newCampaign.Name &&
			campaign.Content == newCampaign.Content &&
			len(campaign.Contacts) == len(newCampaign.Emails)
	})).Return(nil)

	service := Service{Repository: repo}

	id, err := service.Create(newCampaign)

	assert.Nil(err)
	assert.NotEmpty(id)

	repo.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repo := new(repositoryMock)
	repo.On("Save", mock.Anything).Return(errors.New("Internal Server Error"))

	service := Service{Repository: repo}

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("Internal Server Error", err.Error())

	repo.AssertExpectations(t)
	assert.True(errors.Is(err, internalerrors.ErrInternal))
}
