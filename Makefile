#Enviroment variables to config aplication.
#PG*, OAUTH* & etc.
# This may be redefined in user enviroment
export PGHOST ?= 127.0.0.1
export PGPORT ?= 5432
export PGUSER ?= goauthorizer
export PGDATABASE ?= goauthorizerdb
export PGPASSWORD ?= 
export SERVER_LISTEN ?= :8080

# System vars
BIN = goauthorizer
export GOPATH := $(shell pwd)
export PATH := $(PATH):$(GOPATH)/bin


.PHONY : default run build fmt test clean deps dp db_init db_clean


default : build run


db:
	psql


db_clean:
	cat src/goauthorizer/db/drop_schema.sql | psql


db_init: db_clean
	cat src/goauthorizer/db/schema.sql | psql


run: build
	$(BIN)


build:
	go install $(BIN)


fmt:
	go fmt $(BIN)


test:
	go test $(BIN)


clean:
	rm -f "$(BIN)"
	

deps :
	go get github.com/gorilla/mux
	go get github.com/lib/pq
