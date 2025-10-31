package publisher

import (
	"context"
	"time"

	"github.com/chains-lab/countries-svc/internal/events/contracts"
	"github.com/google/uuid"
)

const EventCountryStatusUpdate = "country.status.update"

type UpdateCountryStatus struct {
	CountryID uuid.UUID `json:"country_id"`
	Status    string    `json:"status"`
}

func (s *Service) UpdateCountryStatus(
	ctx context.Context,
	countryID uuid.UUID,
	status string,
) error {
	env := contracts.Envelope[UpdateCountryStatus]{
		Event:     EventCountryStatusUpdate,
		Version:   "1",
		Timestamp: time.Now().UTC(),
		Data: UpdateCountryStatus{
			CountryID: countryID,
			Status:    status,
		},
	}

	return s.publish(
		ctx,
		contracts.TopicCountryV1,
		countryID.String(),
		env,
	)
}
