run: clippy build
	docker compose up

build:
	docker compose build

clippy:
	cargo clippy -- -D warnings

clippy-fix:
	cargo clippy --fix --allow-dirty --allow-staged

infra:
	docker compose up dbpg redis -d

migrate:
	sqlx migrate run

deps:
	cargo install sqlx-cli
