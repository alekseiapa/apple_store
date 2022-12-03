# Apple store

## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop) for building, testing, and deploying containerized applications quickly
- [TablePlus](https://tableplus.com/) as DB UI
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) for DB migrations

### Setup local infrastructure

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

- Run db migration down all versions:

  ```bash
  make migratedown
  ```
