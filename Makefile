run:
	@go run main.go

fmt:
	@go fmt ./...

postgres:
	@docker exec -it some-postgres psql -U postgres

createdb:
	@docker exec -it some-postgres createdb --username=postgres --owner=postgres --encoding=UTF8 inventory_management_system

dropdb:
	@docker exec -it some-postgres dropdb --username=postgres --owner=postgres inventory_management_system

.PHONY: run fmt postgres createdb dropdb