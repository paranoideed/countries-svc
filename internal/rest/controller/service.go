package controller

import (
	"context"

	"github.com/chains-lab/countries-svc/internal/domain/models"
	"github.com/chains-lab/countries-svc/internal/domain/services/country"
	"github.com/chains-lab/logium"
	"github.com/google/uuid"
)

type CountrySvc interface {
	Create(ctx context.Context, name string) (models.Country, error)

	GetByID(ctx context.Context, ID uuid.UUID) (models.Country, error)
	GetByName(ctx context.Context, name string) (models.Country, error)

	Filter(
		ctx context.Context,
		filters country.FilterParams,
		page, size uint64,
	) (models.CountriesCollection, error)

	UpdateStatus(ctx context.Context, countryID uuid.UUID, status string) (models.Country, error)

	Update(ctx context.Context, ID uuid.UUID, params country.UpdateParams) (models.Country, error)
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
