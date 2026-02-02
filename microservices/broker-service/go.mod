module broker-service

go 1.25

require (
	github.com/ViXP/go_sample_projects/microservices/api-view-helpers v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.2.3
	github.com/go-chi/cors v1.2.2
)

require github.com/rabbitmq/amqp091-go v1.10.0

replace github.com/ViXP/go_sample_projects/microservices/api-view-helpers => ../api-view-helpers
