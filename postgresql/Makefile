SHELL := bash

DB_HOST ?= 127.0.0.1
DB_PORT ?= 5432
DB_NAME ?= upperio_tests

TEST_FLAGS ?=

export DB_HOST
export DB_NAME
export DB_PORT

build:
	go build && go install

require-client:
	@if [ -z "$$(which psql)" ]; then \
		echo 'Missing "psql" command. Please install the PostgreSQL client and try again.' && \
		exit 1; \
	fi

generate:
	go generate && \
	go get -d -t -v ./...

reset-db: require-client
	SQL="" && \
	SQL+="DROP DATABASE IF EXISTS $(DB_NAME);" && \
	SQL+="CREATE DATABASE $(DB_NAME) ENCODING 'UTF-8' LC_COLLATE='en_US.UTF-8' LC_CTYPE='en_US.UTF-8' TEMPLATE template0;" && \
	psql -Upostgres -h$(DB_HOST) -p$(DB_PORT) <<< $$SQL; \

test: reset-db generate
	#go test -tags generated -v -race # race: limit on 8192 simultaneously alive goroutines is exceeded, dying
	go test -tags generated -v $(TEST_FLAGS)
