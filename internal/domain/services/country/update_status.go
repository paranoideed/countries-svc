package country

import (
	"context"
	"fmt"
	"time"

	"github.com/chains-lab/countries-svc/internal/domain/enum"
	"github.com/chains-lab/countries-svc/internal/domain/errx"
	"github.com/chains-lab/countries-svc/internal/domain/models"
	"github.com/google/uuid"
)

func (s Service) UpdateStatus(ctx context.Context, countryID uuid.UUID, status string) (models.Country, error) {
	err := enum.CheckCountryStatus(status)
	if err != nil {
		return models.Country{}, errx.ErrorInvalidCountryStatus.Raise(err)
	}

	now := time.Now().UTC()

	cou, err := s.GetByID(ctx, countryID)
	if err != nil {
		return models.Country{}, err
	}

	txErr := s.db.Transaction(ctx, func(ctx context.Context) error {
		switch status {

		case enum.CountryStatusSupported:
			//skip all good wee need only update status
		case enum.CountryStatusUnsupported,
			enum.CountryStatusDeprecated:
			//TODO event
			err = s.eve.UpdateCountryStatus(ctx, countryID, status)
			if err != nil {
				return errx.ErrorInternal.Raise(
					fmt.Errorf("error updating country status, by event: %w", err),
				)
			}
		}

		err = s.db.UpdateCountryStatus(ctx, countryID, status, now)

		cou.Status = status
		cou.UpdatedAt = now

		return nil
	})

	if txErr != nil {
		return models.Country{}, txErr
	}

	return cou, nil
}
