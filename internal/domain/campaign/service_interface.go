package campaign

import "emailn/internal/contract"

type ServiceInterface interface {
    Create(newCampaign contract.NewCampaign) (string, error)
    Get() (interface{}, error)
}
