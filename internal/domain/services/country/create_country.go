package country

import (
	"context"
	"fmt"
	"time"

	"github.com/pariz/gountries"

	"github.com/chains-lab/countries-svc/internal/domain/errx"
	"github.com/chains-lab/countries-svc/internal/domain/models"
)

func (s Service) CreateCountry(ctx context.Context, iso3ID string) (models.Country, error) {
	countryID, err := gountries.New().FindCountryByAlpha(iso3ID)
	if err != nil {
		return models.Country{}, errx.ErrorInvalidCountryISO3ID.Raise(
			fmt.Errorf("invalid country ISO3 ID %s: %w", iso3ID, err),
		)
	}

	_, err = s.GeCountryByID(ctx, countryID.Alpha3)
	if err == nil {
		return models.Country{}, errx.ErrorCountryAlreadyExistsWithThisID.Raise(err)
	}

	var cou models.Country
	var pol models.Policy
	now := time.Now().UTC()

	if err = s.db.Transaction(ctx, func(txCtx context.Context) error {
		cou, pol, err = s.db.CreateCountry(ctx, models.Country{
			ID:        iso3ID,
			CreatedAt: now,
			UpdatedAt: now,
		})

		return nil
	}); err != nil {
		return models.Country{}, errx.ErrorInternal.Raise(
			fmt.Errorf("failed to creating country, cause: %w", err),
		)
	}

	err = s.eve.UpdatedCountryPolicy(ctx, pol)
	if err != nil {
		return models.Country{}, errx.ErrorInternal.Raise(
			fmt.Errorf("failed to publish policy update event, cause: %w", err),
		)
	}

	return cou, nil
}
