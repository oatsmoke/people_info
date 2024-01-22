include .env
export
migrate-up:
	goose -dir ${MIGRATION_PATH} ${DRIVER} ${DB_STRING} up
migrate-down:
	goose -dir ${MIGRATION_PATH} ${DRIVER} ${DB_STRING} down