DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/main.go

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks,users -package api openapi/openapi.yaml > ./internal/web/api.gen.go
	
lint:
	golangci-lint run --color=auto
test:
	go test ./... -v