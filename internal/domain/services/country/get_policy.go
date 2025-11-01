package country

import (
	"context"
	"fmt"

	"github.com/chains-lab/countries-svc/internal/domain/errx"
	"github.com/chains-lab/countries-svc/internal/domain/models"
)

func (s Service) GetPolicyByCountryID(ctx context.Context, countryID string) (models.Policy, error) {
	policy, err := s.db.GetPolicyByCountryID(ctx, countryID)
	if err != nil {
		return models.Policy{}, errx.ErrorInternal.Raise(
			fmt.Errorf("failed to get policy by country id %s, cause: %w", countryID, err),
		)
	}

	if policy.IsNil() {
		return models.Policy{}, errx.ErrorPolicyNotFound.Raise(
			fmt.Errorf("policy not found for country %s", countryID),
		)
	}

	return policy, nil
}
