package controller

import (
	"errors"
	"net/http"

	"github.com/chains-lab/ape"
	"github.com/chains-lab/ape/problems"
	"github.com/chains-lab/countries-svc/internal/domain/errx"
	"github.com/chains-lab/countries-svc/internal/domain/services/country"
	"github.com/chains-lab/countries-svc/internal/rest/requests"
	"github.com/chains-lab/countries-svc/internal/rest/responses"
)

func (a Service) UpdateCountryPolicy(w http.ResponseWriter, r *http.Request) {
	req, err := requests.UpdateCountryPolity(r)
	if err != nil {
		a.log.WithError(err).Error("failed to parse update country policy request")
		ape.RenderErr(w, problems.BadRequest(err)...)

		return
	}

	params := country.PolicyUpdateParams{}
	if req.Data.Attributes.CitiesAllowedStatuses != nil {
		params.CitiesAllowedStatuses = req.Data.Attributes.CitiesAllowedStatuses
	}

	res, err := a.domain.country.UpdatePolicy(r.Context(), req.Data.Id, params)
	if err != nil {
		a.log.WithError(err).Error("failed to update country")
		switch {
		case errors.Is(err, errx.ErrorPolicyNotFound):
			ape.RenderErr(w, problems.NotFound("country policy not found"))
		default:
			ape.RenderErr(w, problems.InternalError())
		}

		return
	}

	ape.Render(w, http.StatusOK, responses.Policy(res))
}
