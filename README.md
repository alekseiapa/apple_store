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
