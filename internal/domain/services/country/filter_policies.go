package country

import (
	"context"

	"github.com/chains-lab/countries-svc/internal/domain/models"
)

type FilterPoliciesParams struct {
	CountryIDs []string
}

func (s Service) FilterPolicies(
	ctx context.Context,
	filters FilterPoliciesParams,
	page, size uint64,
) (models.PoliciesCollection, error) {
	res, err := s.db.FilterPolicies(ctx, filters, page, size)
	if err != nil {
		return models.PoliciesCollection{}, err
	}

	return res, nil
}
