db-create:
	docker exec -i ${DB_CONTAINER} mysql -e "DROP DATABASE IF EXISTS ${DB_NAME};"
	docker exec -i ${DB_CONTAINER} mysql -e "CREATE DATABASE ${DB_NAME} DEFAULT COLLATE utf8mb4_general_ci;"
