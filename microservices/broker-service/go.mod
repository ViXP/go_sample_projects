module broker-service

go 1.25

require (
	github.com/ViXP/go_sample_projects/microservices/api-view-helpers v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.2.3
	github.com/go-chi/cors v1.2.2
)

require (
	github.com/rabbitmq/amqp091-go v1.10.0
	google.golang.org/grpc v1.78.0
	google.golang.org/protobuf v1.36.10
)

require (
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
)

replace github.com/ViXP/go_sample_projects/microservices/api-view-helpers => ../api-view-helpers
