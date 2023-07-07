CREATE TABLE IF NOT EXISTS Products (
  jan BIGINT PRIMARY KEY,
  product_name VARCHAR(255) NOT NULL,
  maker_id INT NOT NULL,
  brand_id INT,
  FOREIGN KEY (maker_id) REFERENCES Makers(id),
  FOREIGN KEY (brand_id) REFERENCES Brands(id)
);