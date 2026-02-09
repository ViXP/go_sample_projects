# Makefile for managing the microservices infrastructure
run:
	@echo "Running the microservices infrastructure"
	docker compose up -d frontend broker auth logger mailer dispatcher
	@echo "The containers started"

brun:
	@echo "Compiling and running the microservices infrastructure"
	docker compose up -d frontend_builder broker_builder auth_builder logger_builder mailer_builder dispatcher_builder
	@echo "The containers started"

build:
	@echo "Building the images for push"
	docker buildx build -f dockerfiles/auth-service.dockerfile -t vixp/auth-service:1.0.0 .
	docker buildx build -f dockerfiles/broker-service.dockerfile -t vixp/broker-service:1.0.0 .
	docker buildx build -f dockerfiles/dispatcher-service.dockerfile -t vixp/dispatcher-service:1.0.0 .
	docker buildx build -f dockerfiles/logger-service.dockerfile -t vixp/logger-service:1.0.0 .
	docker buildx build -f dockerfiles/mailer-service.dockerfile -t vixp/mailer-service:1.0.0 .
	docker buildx build -f dockerfiles/frontend.dockerfile -t vixp/frontend:1.0.1 .
	@echo "The images are built"

push:
	@echo "Pushing the Docker Hub images"
	docker push vixp/auth-service:1.0.0
	docker push vixp/broker-service:1.0.0
	docker push vixp/dispatcher-service:1.0.0
	docker push vixp/logger-service:1.0.0
	docker push vixp/mailer-service:1.0.0
	docker push vixp/frontend:1.0.1
	@echo "All images are pushed"

deploy_swarm:
	@echo "Deploying the services into the local Swarm cluster"
	docker swarm init
	docker stack deploy -c swarm.yml distributed_app
	@echo "Deployed"

db:
	@echo "Starting only the database and message brokers containers"
	docker compose up -d db mongodb mongoexpress pgadmin rabbitmq
	@echo "The database container started"

stop:
	@echo "Stopping the microservices infrastructure"
	docker compose down
	@echo "The containers are stopped"

log:
	@echo "Microservices containers logs:"
	docker compose logs -f
