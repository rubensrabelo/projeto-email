package endpoints

import (
	"fmt"
	"net/http/httptest"
	"testing"
	internalmock "emailn/internal/test/mock"

	"github.com/stretchr/testify/assert"
)

func Test_CampaignsGet_should_return_list(t *testing.T) {
	assert := assert.New(t)
	
	expectedCampaigns := []map[string]interface{}{
		{"id": "1", "name": "Campanha 1"},
		{"id": "2", "name": "Campanha 2"},
	}

	service := new(internalmock.ServiceMock)
	service.On("Get").Return(expectedCampaigns, nil)
	
	handler := Handler{CampaignService: service}
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, err := handler.CampaignGet(rr, req)

	assert.Equal(200, status)
	assert.Nil(err)
	assert.Equal(expectedCampaigns, response)
}

func Test_CampaignsGet_should_return_error(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.ServiceMock)
	service.On("Get").Return(nil, fmt.Errorf("internal error"))
	
	handler := Handler{CampaignService: service}
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignGet(rr, req)

	assert.Equal(200, status)
	assert.NotNil(err)
}
