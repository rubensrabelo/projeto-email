package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(
		newCampaign.Name,
		newCampaign.Content,
		newCampaign.Emails,
	)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}

func (s *Service) Get() (interface{}, error) {
    return s.Repository.Get()
}

func (s *Service) GetBy(id string) (*contract.CampaignResponse, error) {
    campaign, err := s.Repository.GetBy(id)
    if err != nil {
        return nil, internalerrors.ErrInternal
    }

    // Adicione esta verificação de segurança
    if campaign == nil {
        return nil, nil // ou um erro de "not found"
    }

    return &contract.CampaignResponse{
        ID:      campaign.ID,
        Name:    campaign.Name,
        Content: campaign.Content,
        Status:  campaign.Status,
    }, nil
}
