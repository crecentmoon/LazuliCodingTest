CREATE TABLE IF NOT EXISTS ReviewTags (
  id INT AUTO_INCREMENT PRIMARY KEY,
  jan BIGINT,
  tag_from_review VARCHAR(255),
  FOREIGN KEY (jan) REFERENCES Products(jan)
);