package controller

import (
	"context"

	"github.com/chains-lab/countries-svc/internal/domain/models"
	"github.com/chains-lab/countries-svc/internal/domain/services/country"
	"github.com/chains-lab/logium"
)

type CountrySvc interface {
	CreateCountry(ctx context.Context, iso3ID string) (models.Country, error)

	FilterCountry(
		ctx context.Context,
		filters country.FilterCountryParams,
		page, size uint64,
	) (models.CountriesCollection, error)

	FilterPolicies(
		ctx context.Context,
		filters country.FilterPoliciesParams,
		page, size uint64,
	) (models.PoliciesCollection, error)

	GeCountryByID(ctx context.Context, ID string) (models.Country, error)
	GetPolicyByCountryID(ctx context.Context, countryID string) (models.Policy, error)

	UpdatePolicy(ctx context.Context, countryID string, params country.PolicyUpdateParams) (models.Policy, error)
}

type domain struct {
	country CountrySvc
}

type Service struct {
	domain domain
	log    logium.Logger
}

func New(log logium.Logger, country CountrySvc) Service {
	return Service{
		log: log,
		domain: domain{
			country: country,
		},
	}
}
