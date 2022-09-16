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
