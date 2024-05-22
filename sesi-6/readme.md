# LetsGO!

This repo is my notes and mini challenges from Golang For Women 2024 class.

## How to run

Navigate to the **sesi-n** folder, for example:

```bash
cd sesi-6
```

Create the database and tables:

```bash
CREATE DATABASE databasename;

CREATE TABLE products (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL
);

CREATE TABLE variants (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `quantity` INT NOT NULL DEFAULT(0),
  `product_id` INT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  FOREIGN KEY (product_id) REFERENCES products(id)
);
```

Create `.env` file and fill your database information

```bash
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
DB_HOST=
```

Then run

```bash
go run .
```
