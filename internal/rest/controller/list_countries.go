package controller

import (
	"net/http"

	"github.com/chains-lab/ape"
	"github.com/chains-lab/ape/problems"
	"github.com/chains-lab/countries-svc/internal/domain/services/country"
	"github.com/chains-lab/countries-svc/internal/rest/responses"

	"github.com/chains-lab/restkit/pagi"
)

func (a Service) ListCountries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := r.URL.Query()

	filters := country.FilterCountryParams{}

	if sts := q["id"]; len(sts) > 0 {
		filters.IDs = sts
	}

	page, size := pagi.GetPagination(r)

	countries, err := a.domain.country.FilterCountry(ctx, filters, page, size)
	if err != nil {
		a.log.WithError(err).Error("failed to search countries")
		switch {
		default:
			ape.RenderErr(w, problems.InternalError())
		}

		return
	}

	ape.Render(w, http.StatusOK, responses.CountriesCollection(countries))
}
