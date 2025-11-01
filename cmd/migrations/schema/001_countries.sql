-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE countries (
    id   CHAR(3)  PRIMARY KEY,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT chk_countries_id_format CHECK (id ~ '^[A-Z]{3}$')
);

CREATE TABLE country_policies (
    country_id CHAR(3) PRIMARY KEY NOT NULL REFERENCES countries(id) ON DELETE CASCADE,

    cities_allowed_statuses  CITEXT[] NOT NULL DEFAULT ARRAY[]::CITEXT[],

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
);

ALTER TABLE country_policies
ADD CONSTRAINT chk_cities_statuses_lowercase
CHECK (
    cities_statuses_allowed = (
        SELECT ARRAY_AGG(DISTINCT lower(x) ORDER BY lower(x))
        FROM unnest(cities_statuses_allowed) AS x
    )
);

CREATE UNIQUE INDEX ux_country_policies_current
    ON country_policies(country_id)
    WHERE is_current;

CREATE INDEX idx_country_policies_status_allowed
    ON country_policies USING GIN (cities_statuses_allowed);

-- +migrate Down
DROP INDEX IF EXISTS idx_country_policies_status_allowed;
DROP INDEX IF EXISTS ux_country_policies_current;
ALTER TABLE country_policies DROP CONSTRAINT IF EXISTS chk_cities_statuses_lowercase;
DROP TABLE IF EXISTS country_policies CASCADE;
DROP TABLE IF EXISTS countries CASCADE;