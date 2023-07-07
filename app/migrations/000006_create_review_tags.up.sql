CREATE TABLE IF NOT EXISTS ReviewTags (
  id INT PRIMARY KEY,
  jan BIGINT,
  tag_from_review VARCHAR(255),
  FOREIGN KEY (jan) REFERENCES Products(jan)
);