package rest

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/chains-lab/countries-svc/internal"
	"github.com/chains-lab/countries-svc/internal/rest/meta"
	"github.com/chains-lab/logium"
	"github.com/chains-lab/restkit/roles"
	"github.com/go-chi/chi/v5"
)

type Handlers interface {
	CreateCountry(w http.ResponseWriter, r *http.Request)
	ListCountries(w http.ResponseWriter, r *http.Request)
	ListCountryPolicies(w http.ResponseWriter, r *http.Request)
	GetCountry(w http.ResponseWriter, r *http.Request)
	GetCountryPolicy(w http.ResponseWriter, r *http.Request)
	UpdateCountryPolicy(w http.ResponseWriter, r *http.Request)
}

type Middlewares interface {
	Auth(userCtxKey interface{}, skUser string) func(http.Handler) http.Handler
	RoleGrant(userCtxKey interface{}, allowedRoles map[string]bool) func(http.Handler) http.Handler
}

func Run(ctx context.Context, cfg internal.Config, log logium.Logger, m Middlewares, h Handlers) {
	auth := m.Auth(meta.UserCtxKey, cfg.JWT.User.AccessToken.SecretKey)

	sysadmin := m.RoleGrant(meta.UserCtxKey, map[string]bool{
		roles.Admin: true,
	})

	r := chi.NewRouter()

	r.Route("/countries-svc/", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/countries", func(r chi.Router) {
				r.With(auth, sysadmin).Post("/", h.CreateCountry)

				r.Get("/", h.ListCountries)
				r.Get("/policies", h.ListCountryPolicies)

				r.Route("/{country_id}", func(r chi.Router) {
					r.Get("/", h.GetCountry)

					r.Route("/policy", func(r chi.Router) {
						r.Get("/", h.GetCountryPolicy)
						r.With(auth, sysadmin).Put("/", h.UpdateCountryPolicy)
					})
				})
			})
		})
	})

	srv := &http.Server{
		Addr:              cfg.Rest.Port,
		Handler:           r,
		ReadTimeout:       cfg.Rest.Timeouts.Read,
		ReadHeaderTimeout: cfg.Rest.Timeouts.ReadHeader,
		WriteTimeout:      cfg.Rest.Timeouts.Write,
		IdleTimeout:       cfg.Rest.Timeouts.Idle,
	}

	log.Infof("starting REST service on %s", cfg.Rest.Port)

	errCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		} else {
			errCh <- nil
		}
	}()

	select {
	case <-ctx.Done():
		log.Info("shutting down REST service...")
	case err := <-errCh:
		if err != nil {
			log.Errorf("REST server error: %v", err)
		}
	}

	shCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shCtx); err != nil {
		log.Errorf("REST shutdown error: %v", err)
	} else {
		log.Info("REST server stopped")
	}
}
