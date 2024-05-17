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