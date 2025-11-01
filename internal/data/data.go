package data

import (
	"context"
	"database/sql"

	"github.com/chains-lab/countries-svc/internal/data/pgdb"
)

type Database struct {
	sql SqlDB
}

func (d *Database) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.sql.countries.New().Transaction(ctx, fn)
}

type SqlDB struct {
	countries pgdb.CountriesQ
	policies  pgdb.PoliciesQ
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		sql: SqlDB{
			countries: pgdb.NewCountriesQ(db),
			policies:  pgdb.NewPoliciesQ(db),
		},
	}
}
