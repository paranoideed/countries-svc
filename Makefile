DB_URL=postgresql://postgres:postgres@localhost:7000/postgres?sslmode=disable
OPENAPI_GENERATOR := java -jar ~/openapi-generator-cli.jar
CONFIG_FILE := ./config.yaml
API_SRC := ./docs/api.yaml
API_BUNDLED := ./docs/api-bundled.yaml
OUTPUT_DIR := ./docs/web
RESOURCES_DIR := ./resources

generate-models:
	test -d $(RESOURCES_DIR) || mkdir -p $(RESOURCES_DIR)
	test -d $(dir $(API_SRC)) || mkdir -p $(dir $(API_SRC))
	test -d $(dir $(API_BUNDLED)) || mkdir -p $(dir $(API_BUNDLED))
	test -d $(OUTPUT_DIR) || mkdir -p $(OUTPUT_DIR)

	find $(RESOURCES_DIR) -type f ! \( -name "enum_types.go" -o -name "links.go" \) -delete
	swagger-cli bundle $(API_SRC) --outfile $(API_BUNDLED) --type yaml

	$(OPENAPI_GENERATOR) generate \
		-i $(API_BUNDLED) -g go \
		-o $(OUTPUT_DIR) \
		--additional-properties=packageName=resources \
		--import-mappings uuid.UUID=github.com/google/uuid --type-mappings string+uuid=uuid.UUID

	mkdir -p $(RESOURCES_DIR)
	find $(OUTPUT_DIR) -name '*.go' -exec mv {} $(RESOURCES_DIR)/ \;
	find $(RESOURCES_DIR) -type f -name "*_test.go" -delete

sqlc-build:
	KV_VIPER_FILE=$(CONFIG_FILE) sqlc generate

build:
	KV_VIPER_FILE=$(CONFIG_FILE) go build -o ./cmd/countries-svc/main ./cmd/countries-svc/main.go

migrate-up:
	KV_VIPER_FILE=$(CONFIG_FILE) go build -o ./cmd/countries-svc/main ./cmd/countries-svc/main.go
	KV_VIPER_FILE=$(CONFIG_FILE) ./cmd/countries-svc/main migrate up

migrate-down:
	KV_VIPER_FILE=$(CONFIG_FILE) go build -o ./cmd/countries-svc/main ./cmd/countries-svc/main.go
	KV_VIPER_FILE=$(CONFIG_FILE) ./cmd/countries-svc/main migrate down

run-server:
	KV_VIPER_FILE=$(CONFIG_FILE) go build -o ./cmd/countries-svc/main ./cmd/countries-svc/main.go
	KV_VIPER_FILE=$(CONFIG_FILE) ./cmd/countries-svc/main run service

docker-uo:
	docker compose up -d

docker-down:
	docker compose down

docker-rebuild:
	docker compose up -d --build --force-recreate