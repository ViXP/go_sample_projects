# Makefile for managing the microservices infrastructure
run:
	@echo "Running the microservices infrastructure"
	docker compose up -d db frontend broker auth
	@echo "The containers started"

brun:
	@echo "Compiling and running the microservices infrastructure"
	docker compose up -d db frontend_builder broker_builder auth_builder
	@echo "The containers started"

stop:
	@echo "Stopping the microservices infrastructure"
	docker compose down
	@echo "The containers are stopped"

log:
	@echo "Microservices containers logs:"
	docker compose logs -f
