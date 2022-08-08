build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/application

build-docker: build
	docker build -t kombibeer .

run-postgres: 
	docker-compose -f docker-compose-postgres.yml up

run-all: build-docker
	docker-compose -f docker-compose.yml up
	