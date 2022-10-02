.PHONEY: start
start:dc-up
	air -c .air.toml

.PHONEY: dc-up
dc-up:
	docker-compose up -d

.PHONEY: dc-down
dc-down:
	docker-compose down

.PHONEY: init-db-local
init-db-local:dc-up
	docker exec -i go-cms sh /opt/mysql/setup.sh

.PHONEY: init-storage-local
init-storage-local:dc-up
	go run fixtures/storage/main.go

.PHONEY: build-api
build-api:
	go build -o ./api ./cmd/api/main.go

.PHONEY: build-batch
build-batch:
	go build -o ./batch ./cmd/batch/main.go
