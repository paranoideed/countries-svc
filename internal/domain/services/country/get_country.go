package country

import (
	"context"
	"fmt"

	"github.com/chains-lab/countries-svc/internal/domain/errx"
	"github.com/chains-lab/countries-svc/internal/domain/models"
)

func (s Service) GeCountryByID(ctx context.Context, ID string) (models.Country, error) {
	country, err := s.db.GetCountryByID(ctx, ID)
	if err != nil {
		return models.Country{}, errx.ErrorInternal.Raise(
			fmt.Errorf("failed to get country by id %s, cause: %w", ID, err),
		)
	}

	if country.IsNil() {
		return models.Country{}, errx.ErrorCountryNotFound.Raise(
			fmt.Errorf("country not found %s", ID),
		)
	}

	return country, nil
}
