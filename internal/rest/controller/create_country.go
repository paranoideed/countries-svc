package controller

import (
	"errors"
	"net/http"

	"github.com/chains-lab/ape"
	"github.com/chains-lab/ape/problems"
	"github.com/chains-lab/countries-svc/internal/domain/errx"
	"github.com/chains-lab/countries-svc/internal/rest/meta"
	"github.com/chains-lab/countries-svc/internal/rest/requests"
	"github.com/chains-lab/countries-svc/internal/rest/responses"
)

func (a Service) CreateCountry(w http.ResponseWriter, r *http.Request) {
	initiator, err := meta.User(r.Context())
	if err != nil {
		a.log.WithError(err).Error("failed to get user from context")
		ape.RenderErr(w, problems.Unauthorized("failed to get user from context"))

		return
	}

	req, err := requests.CreateCountry(r)
	if err != nil {
		a.log.WithError(err).Error("error creating country")
		ape.RenderErr(w, problems.BadRequest(err)...)

		return
	}

	country, err := a.domain.country.CreateCountry(r.Context(), req.Data.Id)
	if err != nil {
		a.log.WithError(err).Error("error creating country")
		switch {
		case errors.Is(err, errx.ErrorCountryAlreadyExistsWithThisID):
			ape.RenderErr(w, problems.Conflict("country with this id already exists"))
		default:
			ape.RenderErr(w, problems.InternalError())
		}

		return
	}

	a.log.Infof("created country with name %s by user %s", country.ID, initiator.ID)

	ape.Render(w, http.StatusCreated, responses.Country(country))
}
