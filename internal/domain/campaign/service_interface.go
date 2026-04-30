package campaign

import "emailn/internal/contract"

type ServiceInterface interface {
    Create(newCampaign contract.NewCampaign) (string, error)
	GetBy(id string) (*contract.CampaignResponse, error)
	Cancel(id string) error
	Delete(id string) error
}
