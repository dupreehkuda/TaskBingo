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
	go generate ./...

.PHONY: test
test:
	go test ./...
	go test -tags=unit ./internal/usecases/...

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
	docker rm bingo-web
	docker run -d --restart=always --init -p 3000:3000 --name bingo-web ghcr.io/dupreehkuda/bingo-web-prod:latest
