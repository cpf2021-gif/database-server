run:
	@go run main.go

fmt:
	@go fmt ./...

tidy:
	@go mod tidy

shell:
	@docker exec -it some-postgres bash

dumpdb:
	@docker exec -it some-postgres bash -c "pg_dump -U postgres -d inventory_management_system > backup.sql"

restoredb:
	@docker exec -it some-postgres bash -c "psql -U postgres -d inventory_management_system < backup.sql"

postgres:
	@docker exec -it some-postgres psql -U postgres

createdb:
	@docker exec -it some-postgres createdb --username=postgres --owner=postgres --encoding=UTF8 inventory_management_system

dropdb:
	@docker exec -it some-postgres dropdb --username=postgres inventory_management_system

.PHONY: run fmt tidy postgres createdb dropdb shell dumpdb restoredb