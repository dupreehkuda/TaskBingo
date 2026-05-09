.PHONY: compose
compose:
	docker-compose -f docker-compose.local.yml up

.PHONY: compose-down
compose-down:
	docker-compose -f docker-compose.local.yml down --remove-orphans

.PHONY: rebuild
rebuild:
	docker-compose -f docker-compose.local.yml build
	docker-compose -f docker-compose.local.yml down
	docker-compose -f docker-compose.local.yml up -d

.PHONY: fast
fast:
	docker-compose -f docker-compose.local.yml build
	docker-compose -f docker-compose.local.yml down
	docker-compose -f docker-compose.local.yml up

.PHONY: generate
generate:
	go generate ./...

.PHONY: test
test:
	go test ./...
	go test -tags=unit ./internal/usecases/...
