ifneq (,$(wildcard ./build/.env))
    include ./build/.env
    export $(shell sed 's/=.*//' ./dev/.env)
endif

gen:gen_login gen_shop gen_queries

gen_queries:
	@sqlc generate

gen_shop:
		@mkdir -p ./docs
		@mkdir -p generated/
		@protoc	--go_out=generated --go_opt=paths=source_relative \
				--go-grpc_out=generated --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=generated --grpc-gateway_opt=paths=source_relative \
				--openapiv2_out=allow_merge=true,merge_file_name=api_shopV1:docs \
				protos/admin_v1/admin.proto

gen_login:
		@mkdir -p generated/
		@protoc --go_out=generated --go_opt=paths=source_relative \
				--go-grpc_out=generated --go-grpc_opt=paths=source_relative \
				protos/login_v1/login.proto

docker_create_network:
	@docker network create uzum-api || true

docker_infra: docker_create_network
	@docker-compose -f ./build/docker_compose_infra.yml -p uzum up -d

migration_up:
	@goose -dir ./migrations postgres "${DBCONNSTR}" up

migration_down:
	@goose -dir ./migrations postgres "${DBCONNSTR}" down

migration_create:
	@goose -dir ./migrations postgres "${DBCONNSTR}" create default sql

run:
	@go run cmd/main.go