.PHONY: compose
compose:
	docker-compose up

.PHONY: compose-down
compose-down:
	docker-compose down --remove-orphans

.PHONY: rebuild
rebuild:
	docker-compose build
	docker-compose down
	docker-compose up -d

.PHONY: fast
fast:
	docker-compose build
	docker-compose down
	docker-compose up

.PHONY: generate
generate:
	protoc -I api/proto --go_out=plugins=grpc,paths=source_relative:game-service/pkg/api api/proto/serviceData.proto
	protoc -I api/proto --go_out=plugins=grpc,paths=source_relative:user-data-service/pkg/api api/proto/serviceData.proto
	go generate game-service/internal/models/models.go

.PHONY: deploy
deploy:
	docker-compose -f docker-compose.prod.yml pull
	docker-compose -f docker-compose.prod.yml up -d