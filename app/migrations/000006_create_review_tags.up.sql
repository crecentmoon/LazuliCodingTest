CREATE TABLE IF NOT EXISTS ReviewTags (
  tag_id INT PRIMARY KEY,
  jan INT,
  tag_from_review VARCHAR(255),
  FOREIGN KEY (jan) REFERENCES Products(jan)
);