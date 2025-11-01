package controller

import (
	"errors"
	"net/http"

	"github.com/chains-lab/ape"
	"github.com/chains-lab/ape/problems"
	"github.com/chains-lab/countries-svc/internal/domain/errx"
	"github.com/chains-lab/countries-svc/internal/rest/responses"
	"github.com/go-chi/chi/v5"
)

func (a Service) GetCountryPolicy(w http.ResponseWriter, r *http.Request) {
	policy, err := a.domain.country.GetPolicyByCountryID(r.Context(), chi.URLParam(r, "country_id"))
	if err != nil {
		a.log.WithError(err).Error("failed to get country policy")
		switch {
		case errors.Is(err, errx.ErrorPolicyNotFound):
			ape.RenderErr(w, problems.NotFound("country policy not found"))
		default:
			ape.RenderErr(w, problems.InternalError())
		}

		return
	}

	ape.Render(w, http.StatusOK, responses.Policy(policy))
}
