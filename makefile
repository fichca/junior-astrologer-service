export APOD_CLIENT_API_KEY=
export APOD_CLIENT_BASE_URL=
export HTTP_PORT=
export MINIO_BUCKET=
export POSTGRE_DRIVER=
export POSTGRE_SSL_MODE=

export MINIO_ROOT_USER
export MINIO_ROOT_PASSWORD=

export POSTGRES_DB
export POSTGRES_USER=
export POSTGRES_PASSWORD=

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

test:
	go test -v ./...