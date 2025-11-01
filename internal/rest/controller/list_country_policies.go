package controller

import (
	"net/http"

	"github.com/chains-lab/ape"
	"github.com/chains-lab/ape/problems"
	"github.com/chains-lab/countries-svc/internal/domain/services/country"
	"github.com/chains-lab/countries-svc/internal/rest/responses"

	"github.com/chains-lab/restkit/pagi"
)

func (a Service) ListCountryPolicies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := r.URL.Query()

	filters := country.FilterPoliciesParams{}

	if sts := q["country_id"]; len(sts) > 0 {
		filters.CountryIDs = sts
	}

	page, size := pagi.GetPagination(r)

	policies, err := a.domain.country.FilterPolicies(ctx, filters, page, size)
	if err != nil {
		a.log.WithError(err).Error("failed to search policies")
		switch {
		default:
			ape.RenderErr(w, problems.InternalError())
		}

		return
	}

	ape.Render(w, http.StatusOK, responses.PoliciesCollection(policies))
}
