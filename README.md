goauthorizer
============

Web service to handle authorisation and registration of other web sites


Installation
------------

```
# create postgres db
sudo -u postgres createuser goauthorizer
sudo -u postgres createdb goauthorizerdb -O goauthorizer

# get code
git clone git@github.com:vladimirbright/goauthorizer.git
cd goauthorizer

# create schema
make db_init

# Install dependencys
make deps

# compile and run server on 8080 port
make

```


Configuration
-------------

Server configured by environment variables. See full list in start of Makefile.

Defaults:
```
export PGHOST ?= 127.0.0.1
export PGPORT ?= 5432
export PGUSER ?= goauthorizer
export PGDATABASE ?= goauthorizerdb
export PGPASSWORD ?= 
export SERVER_LISTEN ?= :8080
```

Example:
```
cd goauthorizer
# run on another port
SERVER_LISTEN=:8888 make
```
