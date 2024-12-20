start_db:
	docker-compose -f database/postgres-compose.yml up -d
stop_db:
	docker-compose -f database/postgres-compose.yml down