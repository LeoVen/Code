run: check build
	docker compose up

infra:
	docker compose up dbpg redis -d

check:
	cargo clippy -- -D warnings

check-fix:
	cargo clippy --fix --allow-dirty

build:
	docker compose build

.PHONY: check build run
