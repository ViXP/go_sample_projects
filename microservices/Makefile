# Makefile for managing the microservices infrastructure
run:
	@echo "Running the microservices infrastructure"
	docker compose up -d db mongodb mailhog frontend broker auth logger mailer
	@echo "The containers started"

brun:
	@echo "Compiling and running the microservices infrastructure"
	docker compose up -d db mongodb mailhog frontend_builder broker_builder auth_builder logger_builder mailer_builder
	@echo "The containers started"

db:
	@echo "Starting only the database containers"
	docker compose up -d db mongodb mongoexpress pgadmin
	@echo "The database container started"

stop:
	@echo "Stopping the microservices infrastructure"
	docker compose down
	@echo "The containers are stopped"

log:
	@echo "Microservices containers logs:"
	docker compose logs -f
