db-init:
	docker exec -i ${DB_CONTAINER} mysql < db.sql
