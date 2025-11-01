package publisher

import (
	"context"
	"time"

	"github.com/chains-lab/countries-svc/internal/domain/models"
	"github.com/chains-lab/countries-svc/internal/events/contracts"
)

const EventCountryPoliciesUpdated = "country.policy.updated"

type UpdatedCountryPolicy = models.Policy

func (s *Service) UpdatedCountryPolicy(
	ctx context.Context,
	policy models.Policy,
) error {
	env := contracts.Envelope[UpdatedCountryPolicy]{
		Event:     EventCountryPoliciesUpdated,
		Version:   "1",
		Timestamp: time.Now().UTC(),
		Data:      policy,
	}

	return s.publish(
		ctx,
		contracts.TopicCountryV1,
		policy.CountryID,
		env,
	)
}
