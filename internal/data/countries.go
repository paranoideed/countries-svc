package data

import (
	"context"
	"database/sql"
	"errors"

	"github.com/chains-lab/countries-svc/internal/data/pgdb"
	"github.com/chains-lab/countries-svc/internal/domain/models"
	"github.com/chains-lab/countries-svc/internal/domain/services/country"
	"github.com/chains-lab/restkit/pagi"
)

func (d *Database) CreateCountry(ctx context.Context, country models.Country) (models.Country, models.Policy, error) {
	err := d.sql.countries.New().Insert(ctx, pgdb.Country{
		ID:        country.ID,
		CreatedAt: country.CreatedAt,
		UpdatedAt: country.UpdatedAt,
	})
	if err != nil {
		return models.Country{}, models.Policy{}, err
	}

	err = d.sql.policies.New().Insert(ctx, pgdb.Policy{
		CountryID:             country.ID,
		CitiesAllowedStatuses: []string{},
		CreatedAt:             country.CreatedAt,
		UpdatedAt:             country.UpdatedAt,
	})
	if err != nil {
		return models.Country{}, models.Policy{}, err
	}

	return country, models.Policy{
		CountryID:             country.ID,
		CitiesAllowedStatuses: []string{},
		CreatedAt:             country.CreatedAt,
		UpdatedAt:             country.UpdatedAt,
	}, nil
}

func (d *Database) GetCountryByID(ctx context.Context, ID string) (models.Country, error) {
	c, err := d.sql.countries.New().FilterID(ID).Get(ctx)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return models.Country{}, nil
	case err != nil:
		return models.Country{}, err
	}

	return countrySchemaToModel(c), err
}

func (d *Database) FilterCountries(
	ctx context.Context,
	filters country.FilterCountryParams,
	page, size uint64,
) (models.CountriesCollection, error) {
	limit, offset := pagi.PagConvert(page, size)

	query := d.sql.countries.New()

	if filters.IDs != nil {
		query = query.FilterID(filters.IDs...)
	}

	rows, err := query.Page(limit, offset).Select(ctx)
	if err != nil {
		return models.CountriesCollection{}, err
	}

	total, err := query.Count(ctx)
	if err != nil {
		return models.CountriesCollection{}, err
	}

	countries := make([]models.Country, 0, len(rows))
	for _, row := range rows {
		countries = append(countries, countrySchemaToModel(row))
	}

	return models.CountriesCollection{
		Data:  countries,
		Page:  page,
		Size:  size,
		Total: total,
	}, nil
}

func countrySchemaToModel(row pgdb.Country) models.Country {
	return models.Country{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
