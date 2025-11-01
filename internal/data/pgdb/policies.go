package pgdb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

const policiesTable = "country_policies"

type Policy struct {
	CountryID             string    `db:"country_id"`
	CitiesAllowedStatuses []string  `db:"cities_allowed_statuses"`
	CreatedAt             time.Time `db:"created_at"`
	UpdatedAt             time.Time `db:"updated_at"`
}

type PoliciesQ struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewPoliciesQ(db *sql.DB) PoliciesQ {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return PoliciesQ{
		db:       db,
		selector: builder.Select("*").From(policiesTable),
		inserter: builder.Insert(policiesTable),
		updater:  builder.Update(policiesTable),
		deleter:  builder.Delete(policiesTable),
		counter:  builder.Select("COUNT(*) AS count").From(policiesTable),
	}
}

func (q PoliciesQ) New() PoliciesQ {
	return NewPoliciesQ(q.db)
}

func (q PoliciesQ) Insert(ctx context.Context, input Policy) error {
	values := map[string]interface{}{
		"country_id":              input.CountryID,
		"cities_allowed_statuses": input.CitiesAllowedStatuses,
		"created_at":              input.CreatedAt,
		"updated_at":              input.UpdatedAt,
	}

	query, args, err := q.inserter.SetMap(values).ToSql()
	if err != nil {
		return err
	}

	if tx, ok := TxFromCtx(ctx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = q.db.ExecContext(ctx, query, args...)
	}

	return err
}

func (q PoliciesQ) Get(ctx context.Context) (Policy, error) {
	query, args, err := q.selector.ToSql()
	if err != nil {
		return Policy{}, err
	}

	var policy Policy
	var row *sql.Row
	if tx, ok := TxFromCtx(ctx); ok {
		row = tx.QueryRowContext(ctx, query, args...)
	} else {
		row = q.db.QueryRowContext(ctx, query, args...)
	}
	err = row.Scan(
		&policy.CountryID,
		pq.Array(&policy.CitiesAllowedStatuses),
		&policy.CreatedAt,
		&policy.UpdatedAt,
	)

	return policy, err
}

func (q PoliciesQ) Select(ctx context.Context) ([]Policy, error) {
	query, args, err := q.selector.ToSql()
	if err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if tx, ok := TxFromCtx(ctx); ok {
		rows, err = tx.QueryContext(ctx, query, args...)
	} else {
		rows, err = q.db.QueryContext(ctx, query, args...)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var policies []Policy
	for rows.Next() {
		var policy Policy
		err := rows.Scan(
			&policy.CountryID,
			pq.Array(&policy.CitiesAllowedStatuses),
			&policy.CreatedAt,
			&policy.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		policies = append(policies, policy)
	}

	return policies, nil
}

func (q PoliciesQ) Update(ctx context.Context, updatedAt time.Time) error {
	q.updater = q.updater.Set("updated_at", updatedAt)

	query, args, err := q.updater.ToSql()
	if err != nil {
		return fmt.Errorf("building update query for %s: %w", policiesTable, err)
	}

	if tx, ok := TxFromCtx(ctx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = q.db.ExecContext(ctx, query, args...)
	}
	return err
}

func (q PoliciesQ) UpdateCitiesManagement(cities []string) PoliciesQ {
	q.updater = q.updater.Set("cities_allowed_statuses", sq.Expr("?::citext[]", pq.Array(cities)))
	return q
}

func (q PoliciesQ) Delete(ctx context.Context) error {
	query, args, err := q.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building deleter query for table: %s: %w", policiesTable, err)
	}

	if tx, ok := TxFromCtx(ctx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = q.db.ExecContext(ctx, query, args...)
	}

	return err
}

func (q PoliciesQ) FilterCountryID(countryID ...string) PoliciesQ {
	q.selector = q.selector.Where(sq.Eq{"country_id": countryID})
	q.updater = q.updater.Where(sq.Eq{"country_id": countryID})
	q.deleter = q.deleter.Where(sq.Eq{"country_id": countryID})
	q.counter = q.counter.Where(sq.Eq{"country_id": countryID})
	return q
}

func (q PoliciesQ) FilterCitiesHasAll(statuses []string) PoliciesQ {
	q.selector = q.selector.Where(sq.Expr("cities_allowed_statuses @> ?::citext[]", pq.Array(statuses)))
	q.counter = q.counter.Where(sq.Expr("cities_allowed_statuses @> ?::citext[]", pq.Array(statuses)))
	q.updater = q.updater.Where(sq.Expr("cities_allowed_statuses @> ?::citext[]", pq.Array(statuses)))
	q.deleter = q.deleter.Where(sq.Expr("cities_allowed_statuses @> ?::citext[]", pq.Array(statuses)))
	return q
}

func (q PoliciesQ) FilterCitiesHasAny(statuses []string) PoliciesQ {
	q.selector = q.selector.Where(sq.Expr("cities_allowed_statuses && ?::citext[]", pq.Array(statuses)))
	q.counter = q.counter.Where(sq.Expr("cities_allowed_statuses && ?::citext[]", pq.Array(statuses)))
	q.updater = q.updater.Where(sq.Expr("cities_allowed_statuses && ?::citext[]", pq.Array(statuses)))
	q.deleter = q.deleter.Where(sq.Expr("cities_allowed_statuses && ?::citext[]", pq.Array(statuses)))
	return q
}

func (q PoliciesQ) Count(ctx context.Context) (uint64, error) {
	query, args, err := q.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for table: %s: %w", policiesTable, err)
	}

	var row *sql.Row
	if tx, ok := TxFromCtx(ctx); ok {
		row = tx.QueryRowContext(ctx, query, args...)
	} else {
		row = q.db.QueryRowContext(ctx, query, args...)
	}

	var count uint64
	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("scanning count result for table: %s: %w", policiesTable, err)
	}

	return count, nil
}

func (q PoliciesQ) Page(limit, offset uint64) PoliciesQ {
	q.selector = q.selector.Limit(limit).Offset(offset)
	return q
}
