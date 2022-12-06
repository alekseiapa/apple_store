# Apple store

The service is a straightforward store. It will provide APIs for the frontend to perform the following tasks:

1. Create/Read/Update/Delete Products
2. Create/Read/Update/Delete Users
3. Create/Read/Delete Orders
4. Login Users

### Documentation

- [API](https://documenter.getpostman.com/view/15139360/2s8YzP1jcW)
- [DB schema](./apple_store_db.png)
- [DB queries](./db/migration/000001_init_schema.up.sql)

### Tools that were used during development

- [Docker desktop](https://www.docker.com/products/docker-desktop) for building, testing, and deploying containerized applications quickly
- [TablePlus](https://tableplus.com/) as DB UI
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) for DB migrations

# Dev environment

## Run in docker

- Create docker network

  ```bash
  make docker-network
  ```

- Start project in docker compose:

  ```bash
  make docker_compose
  ```

## Run locally without docker

- Start postgres container:

  ```bash
  make postgres
  ```

- Create apple_store database:

  ```bash
  make createdb
  ```

- Run db migration up all versions:

  ```bash
  make migrateup

  ```

- Start api server:

  ```bash
  make server
  ```

## Other commands

- Run all tests:

  ```bash
  make test
  ```

- Drop apple_store database:

  ```bash
  make dropdb
  ```

- Run db migration down all versions:

  ```bash
  make migratedown1ver
  ```

- Run db migration up 1 version:

  ```bash
  make migrateup1ver
  ```

- Run db migration down 1 version:

  ```bash
  make migratedown1
  ```
