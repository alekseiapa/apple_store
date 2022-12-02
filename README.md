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

What I am using:

- Docker for building, testing, and deploying containerized applications quickly https://docs.docker.com
- TablePlus as DB UI https://tableplus.com/download
- Golang-migrate for DB migrations: https://github.com/golang-migrate/migrate
- SQLC for DB interaction https://docs.sqlc.dev/en/latest
