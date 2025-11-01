package responses

import (
	"github.com/chains-lab/countries-svc/internal/domain/models"
	"github.com/chains-lab/countries-svc/resources"
)

func Policy(model models.Policy) resources.CountryPolicy {
	return resources.CountryPolicy{
		Data: resources.CountryPolicyData{
			Id:   model.CountryID,
			Type: resources.PolicyType,
			Attributes: resources.CountryPolicyAttributes{
				CitiesAllowedStatuses: model.CitiesAllowedStatuses,
				CreatedAt:             model.CreatedAt,
				UpdatedAt:             model.UpdatedAt,
			},
		},
	}
}

func PoliciesCollection(modelsCollection models.PoliciesCollection) resources.CountryPoliciesCollection {
	resp := resources.CountryPoliciesCollection{
		Data: make([]resources.CountryPolicyData, 0, len(modelsCollection.Data)),
		Links: resources.PaginationData{
			PageNumber: int64(modelsCollection.Page),
			PageSize:   int64(modelsCollection.Size),
			TotalItems: int64(modelsCollection.Total),
		},
	}

	for _, model := range modelsCollection.Data {
		policy := Policy(model).Data

		resp.Data = append(resp.Data, policy)
	}

	return resp
}
