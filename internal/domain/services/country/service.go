package country

import (
	"context"
	"time"

	"github.com/chains-lab/countries-svc/internal/domain/models"
)

type Service struct {
	db  database
	eve EventWriter
}

func NewService(db database, eve EventWriter) Service {
	return Service{
		db:  db,
		eve: eve,
	}
}

type database interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error

	CreateCountry(ctx context.Context, country models.Country) (models.Country, models.Policy, error)

	GetCountryByID(ctx context.Context, ID string) (models.Country, error)

	GetPolicyByCountryID(ctx context.Context, countryID string) (models.Policy, error)

	FilterCountries(
		ctx context.Context,
		filters FilterCountryParams,
		page, size uint64,
	) (models.CountriesCollection, error)
	FilterPolicies(
		ctx context.Context,
		filters FilterPoliciesParams,
		page, size uint64,
	) (models.PoliciesCollection, error)

	UpdatePolicy(
		ctx context.Context,
		countryID string,
		params PolicyUpdateParams,
		updatedAt time.Time,
	) error
}

type EventWriter interface {
	UpdatedCountryPolicy(
		ctx context.Context,
		policy models.Policy,
	) error
}
