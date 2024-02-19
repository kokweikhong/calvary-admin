docker-up:
	@echo "Starting docker containers"
	@docker compose up -d
	@echo "Docker containers started successfully"

docker-down:
	@echo "Stopping docker containers"
	@docker compose down
	@echo "Docker containers stopped successfully"

docker-restart:
	@echo "Restarting docker containers"
	@docker compose restart
	@echo "Docker containers restarted successfully"

docker-ps:
	@echo "Listing docker containers"
	@docker compose ps
	@echo "Docker containers listed successfully"

docker-logs:
	@echo "Showing docker logs"
	@docker compose logs -f
	@echo "Docker logs shown successfully"

docker-prune:
	@echo "Pruning docker containers"
	@docker system prune -a
	@echo "Docker containers pruned successfully"

docker-up-postgres:
	@echo "Starting postgres container"
	@docker compose up -d postgres
	@echo "Postgres container started successfully"

POSTGRES_URL=postgres://calvaryadmin:calvary@admin@999@localhost:5432/calvary_admin?sslmode=disable

migrate-create:
	@echo "Creating migration file"
	@migrate create -ext sql -dir db/migrations -seq $(name)
	@echo "Migration file created"

migrate-up:
	@echo "Running migrations"
	@migrate -database $(POSTGRES_URL) -path db/migrations up
	@echo "Migrations ran successfully"

migrate-down:
	@echo "Rolling back migrations"
	@migrate -database $(POSTGRES_URL) -path db/migrations down
	@echo "Migrations rolled back successfully"

migrate-force:
	@echo "Forcing migrations"
	@migrate -database $(POSTGRES_URL) -path db/migrations force $(version)
	@echo "Migrations forced successfully"

migrate-version:
	@echo "Getting migration version"
	@migrate -database $(POSTGRES_URL) -path db/migrations version
	@echo "Migration version retrieved successfully"

migrate-drop:
	@echo "Dropping migrations"
	@migrate -database $(POSTGRES_URL) -path db/migrations drop
	@echo "Migrations dropped successfully"

migrate-reset:
	@echo "Resetting migrations"
	@migrate -database $(POSTGRES_URL) -path db/migrations reset
	@echo "Migrations reset successfully"

migrate-status:
	@echo "Getting migration status"
	@migrate -database $(POSTGRES_URL) -path db/migrations status
	@echo "Migration status retrieved successfully"
