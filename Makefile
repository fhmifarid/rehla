.PHONY: api worker migrate-up migrate-status sqlc test test-backend test-admin lint build dev-up dev-down

api:
	cd backend && go run ./cmd/api

worker:
	cd backend && go run ./cmd/worker

migrate-up:
	cd backend && go run ./cmd/migrate up

migrate-status:
	cd backend && go run ./cmd/migrate status

sqlc:
	cd backend && go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.30.0 generate

test: test-backend test-admin

test-backend:
	cd backend && go test -race ./...

test-admin:
	npm run test --workspaces --if-present

lint:
	cd backend && go vet ./...
	npm run lint --workspaces --if-present
	npm run typecheck --workspaces --if-present

build:
	cd backend && go build ./cmd/api ./cmd/worker ./cmd/migrate
	npm run build --workspaces --if-present

dev-up:
	docker compose up -d

dev-down:
	docker compose down
