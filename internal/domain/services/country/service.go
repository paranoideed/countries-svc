package country

import (
	"context"
	"time"

	"github.com/chains-lab/countries-svc/internal/domain/models"
	"github.com/google/uuid"
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

	CreateCountry(ctx context.Context, country models.Country) (models.Country, error)

	GetCountryByID(ctx context.Context, ID uuid.UUID) (models.Country, error)
	GetCountryByName(ctx context.Context, name string) (models.Country, error)

	FilterCountries(
		ctx context.Context,
		filters FilterParams,
		page, size uint64,
	) (models.CountriesCollection, error)

	UpdateCountry(
		ctx context.Context,
		countryID uuid.UUID,
		params UpdateParams,
		updatedAt time.Time,
	) error

	UpdateCountryStatus(ctx context.Context, countryID uuid.UUID, status string, updatedAt time.Time) error
}

type EventWriter interface {
	UpdateCountryStatus(
		ctx context.Context,
		countryID uuid.UUID,
		status string,
	) error
}
