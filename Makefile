.PHONY: run migrate-up migrate-down

run:
	go run cmd/simple-shop/main.go

migrate-create:
	migrate create -ext sql -dir internal/databases/schemas -seq $(action)

migrate-up:
	migrate -path internal/databases/schemas -database "postgresql://root:secret@localhost:5432/simple_shop?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/databases/schemas -database "postgresql://root:secret@localhost:5432/simple_shop?sslmode=disable" -verbose down

migrate-force:
	migrate -path internal/databases/schemas -database "postgresql://root:secret@localhost:5432/simple_shop?sslmode=disable" force $(version)

help:
	@echo "run: run main.go"