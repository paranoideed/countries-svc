package country

import (
	"context"
	"fmt"
	"time"

	"github.com/chains-lab/countries-svc/internal/domain/errx"
	"github.com/chains-lab/countries-svc/internal/domain/models"
)

type PolicyUpdateParams struct {
	CitiesAllowedStatuses []string
}

func (s Service) UpdatePolicy(ctx context.Context, countryID string, params PolicyUpdateParams) (models.Policy, error) {
	pol, err := s.GetPolicyByCountryID(ctx, countryID)
	if err != nil {
		return models.Policy{}, err
	}

	if params.CitiesAllowedStatuses != nil {
		pol.CitiesAllowedStatuses = params.CitiesAllowedStatuses
	}

	now := time.Now().UTC()
	err = s.db.UpdatePolicy(ctx, countryID, params, now)
	if err != nil {
		return models.Policy{}, errx.ErrorInternal.Raise(
			fmt.Errorf("failed to update policy, cause: %w", err),
		)
	}

	err = s.eve.UpdatedCountryPolicy(ctx, pol)
	if err != nil {
		return models.Policy{}, errx.ErrorInternal.Raise(
			fmt.Errorf("failed to publish policy update event, cause: %w", err),
		)
	}

	pol.UpdatedAt = now

	return pol, nil
}
