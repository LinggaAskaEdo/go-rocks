.PHONY: build
build: swag-build
	@go mod tidy && \
		go generate ./src/cmd && \
		go build -o ./build/app ./src/cmd

.PHONY: run
run: swag-build build
	@./build/app

.PHONY: docker-build
docker-build: swag-build build
	@docker build --tag go-rocks .

.PHONY: docker-run
docker-run: 
	@docker run --name go-rocks -e SECRETKEY="$(secretkey)" -p 8181:8181 go-rocks:latest

.PHONY: compose-build
compose-build:
	@docker compose build

.PHONY: compose-up
compose-up: compose-build
	@docker compose up

.PHONY: compose-down
compose-down:
	@docker compose down		

.PHONY: db-start
db-start: 
	@docker start mysql-docker postgres-docker redis-docker

.PHONY: db-stop
db-stop: 
	@docker stop mysql-docker postgres-docker redis-docker

.PHONY: queue-start
queue-start: 
	@docker start rabbitmq-docker

.PHONY: queue-stop
queue-stop: 
	@docker stop rabbitmq-docker

.PHONY: keycloak-start
keycloak-start: 
	@docker start keycloak-docker

.PHONY: keycloak-stop
keycloak-stop: 
	@docker stop keycloak-docker

.PHONY: start-all
start-all: keycloak-start db-start queue-start

.PHONY: stop-all
stop-all: keycloak-stop db-stop queue-stop

.PHONY: swag-install
swag-install:
	@go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: swag-build
swag-build:
	@rm -rf ./docs/
	@swag fmt
	@swag init --quiet --generalInfo ./src/cmd/app.go --parseGoList false

.PHONY: cert-install
cert-install:
	@sudo apt install openssl

.PHONY: cert-create
cert-create:
	@if ! ls -AU "./etc/cert/" | read _; then \
		openssl genrsa -out ./etc/cert/id_rsa 4096 && openssl rsa -in ./etc/cert/id_rsa -pubout -out ./etc/cert/id_rsa.pub; \
	else \
		echo "Directory is not empty !!!"; \
	fi

# jet -source=mysql -dsn="user:pass@tcp(localhost:3306)/dbname" -path=./.gen
.PHONY: jet
jet: build
	@rm -rf ./.gen/*; \
		echo "Please enter source: "; \
		read source; \
		echo "Please enter username: "; \
		read username; \
		echo "Please enter password: "; \
		stty -echo; \
		read passwd; \
		stty echo; \
		echo "Please enter host: "; \
		read host; \
		echo "Please enter port: "; \
		read port; \
		echo "Please enter database: "; \
		read database; \
		echo "Please enter path: "; \
		read genpath; \
		clear; \
		echo "Source: $$source"; \
		echo "Username: $$username"; \
		echo "Password: $$passwd"; \
		echo "Host: $$host"; \
		echo "Port: $$port"; \
		echo "Database: $$database"; \
		echo "Path: $$genpath"; \
		jet -source=$$source -dsn="$$username:$$passwd@tcp($$host:$$port)/$$database" -path=$$genpath