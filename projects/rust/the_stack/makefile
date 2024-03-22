up: clippy build
	docker compose up -d
	sleep 1
	sqlx migrate run
	docker compose logs -f

build:
	docker compose build

clippy:
	cargo clippy --fix --allow-dirty --allow-staged -- -D warnings

down:
	docker compose down

infra:
	docker compose up dbpg redis -d

migrate:
	sqlx migrate run

deps:
	cargo install sqlx-cli
