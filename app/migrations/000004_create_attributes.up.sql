CREATE TABLE IF NOT EXISTS Attributes (
  id INT PRIMARY KEY,
  jan INT,
  attribute_data JSON,
  FOREIGN KEY (jan) REFERENCES Products(jan)
);
