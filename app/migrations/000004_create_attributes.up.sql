CREATE TABLE IF NOT EXISTS Attributes (
  attribute_id INT PRIMARY KEY,
  jan INT,
  attribute_data JSON,
  FOREIGN KEY (jan) REFERENCES Products(jan)
);
