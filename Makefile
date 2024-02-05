include .env

run-api:
	go run cmd/api/main.go

docker-build:	
	docker build -t go-example:deploy -f docker/Dockerfile . 
docker-up:
	docker-compose up -d
docker-down:
	docker-compose down

migrate-create:
	migrate create -ext sql -dir migrations/postgres -seq ${TABLE_NAME}
migrate-up:	
	migrate --verbose -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DATABASE}?sslmode=disable" -path migrations/postgres up
migrate-down:	
	migrate --verbose -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DATABASE}?sslmode=disable" -path migrations/postgres down
migrate-force:	
	migrate --verbose -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DATABASE}?sslmode=disable" -path migrations/postgres force ${V}

linter:
	@echo Starting linters
	golangci-lint run ./...

mock:
	mockgen -source=${SUR} -destination=${DES} -package=mock

test: 
	go test -coverpkg=./... ./... 