.PHONY: compose
compose:
	docker-compose up

.PHONY: compose-down
compose-down:
	docker-compose down --remove-orphans

.PHONY: rebuild
rebuild:
	docker-compose down --remove-orphans
	docker-compose build

.PHONY: fast
fast:
	docker-compose build
	docker-compose down --remove-orphans
	docker-compose up