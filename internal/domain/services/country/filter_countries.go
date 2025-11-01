package country

import (
	"context"
	"fmt"

	"github.com/chains-lab/countries-svc/internal/domain/errx"
	"github.com/chains-lab/countries-svc/internal/domain/models"
)

type FilterCountryParams struct {
	IDs []string
}

func (s Service) FilterCountry(
	ctx context.Context,
	filters FilterCountryParams,
	page, size uint64,
) (models.CountriesCollection, error) {
	res, err := s.db.FilterCountries(ctx, filters, page, size)
	if err != nil {
		return models.CountriesCollection{}, errx.ErrorInternal.Raise(
			fmt.Errorf("failed to filter countries, cause: %w", err),
		)
	}

	return res, nil
}
