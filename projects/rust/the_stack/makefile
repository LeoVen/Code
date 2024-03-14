run: check build
	docker compose up

check:
	cargo clippy

check-fix:
	cargo clippy --fix --allow-dirty

build:
	docker compose build

.PHONY: check build run
