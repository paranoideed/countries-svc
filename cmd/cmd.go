package cmd

import (
	"context"
	"database/sql"
	"sync"

	"github.com/chains-lab/countries-svc/internal"
	"github.com/chains-lab/countries-svc/internal/data"
	"github.com/chains-lab/countries-svc/internal/domain/services/country"
	"github.com/chains-lab/countries-svc/internal/rest"
	"github.com/chains-lab/countries-svc/internal/rest/controller"
	"github.com/chains-lab/countries-svc/internal/rest/middlewares"

	"github.com/chains-lab/logium"
)

func StartServices(ctx context.Context, cfg internal.Config, log logium.Logger, wg *sync.WaitGroup) {
	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	pg, err := sql.Open("postgres", cfg.Database.SQL.URL)
	if err != nil {
		log.Fatal("failed to connect to database", "error", err)
	}

	database := data.NewDatabase(pg)

	countrySvc := country.NewService(database)

	ctrl := controller.New(log, countrySvc)
	mdlv := middlewares.New(log)

	run(func() { rest.Run(ctx, cfg, log, mdlv, ctrl) })
}
