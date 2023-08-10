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
	docker-compose -f docker-compose.prod.yml down
	docker-compose -f docker-compose.prod.yml up -d

.PHONY: web-deploy
web-deploy:
	./move_env.sh
	docker pull ghcr.io/dupreehkuda/bingo-web-prod:latest
	docker kill bingo-web
	docker run -d --restart=always --init -p 3000:3000 --name bingo-web ghcr.io/dupreehkuda/bingo-web-prod:latest