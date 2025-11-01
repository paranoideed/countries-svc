package data

import (
	"context"
	"time"

	"github.com/chains-lab/countries-svc/internal/data/pgdb"
	"github.com/chains-lab/countries-svc/internal/domain/models"
	"github.com/chains-lab/countries-svc/internal/domain/services/country"
	"github.com/chains-lab/restkit/pagi"
)

func (d *Database) GetPolicyByCountryID(ctx context.Context, countryID string) (models.Policy, error) {
	policySchema, err := d.sql.policies.New().FilterCountryID(countryID).Get(ctx)
	if err != nil {
		return models.Policy{}, err
	}

	return policySchemaToModel(policySchema), nil
}

func (d *Database) FilterPolicies(
	ctx context.Context,
	filters country.FilterPoliciesParams,
	page, size uint64,
) (models.PoliciesCollection, error) {
	limit, offset := pagi.PagConvert(page, size)

	query := d.sql.policies.New()

	if filters.CountryIDs != nil {
		query = query.FilterCountryID(filters.CountryIDs...)
	}

	total, err := query.Count(ctx)
	if err != nil {
		return models.PoliciesCollection{}, err
	}

	rows, err := query.Page(limit, offset).Select(ctx)
	if err != nil {
		return models.PoliciesCollection{}, err
	}

	policies := make([]models.Policy, len(rows))
	for i, policySchema := range rows {
		policies[i] = policySchemaToModel(policySchema)
	}

	return models.PoliciesCollection{
		Data:  policies,
		Page:  page,
		Size:  size,
		Total: total,
	}, nil
}

func (d *Database) UpdatePolicy(
	ctx context.Context,
	countryID string,
	params country.PolicyUpdateParams,
	updatedAt time.Time,
) error {
	q := d.sql.policies.New().FilterCountryID(countryID)

	upNotEmpty := false

	if params.CitiesAllowedStatuses != nil {
		q = q.UpdateCitiesManagement(params.CitiesAllowedStatuses)
		upNotEmpty = true
	}

	if !upNotEmpty {
		return nil
	}

	err := q.Update(ctx, updatedAt)
	if err != nil {
		return err
	}

	return nil
}

func policySchemaToModel(policySchema pgdb.Policy) models.Policy {
	return models.Policy{
		CountryID:             policySchema.CountryID,
		CitiesAllowedStatuses: policySchema.CitiesAllowedStatuses,
		CreatedAt:             policySchema.CreatedAt,
		UpdatedAt:             policySchema.UpdatedAt,
	}
}
