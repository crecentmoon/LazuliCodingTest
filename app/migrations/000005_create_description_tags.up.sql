CREATE TABLE IF NOT EXISTS DescriptionTags (
  id INT AUTO_INCREMENT PRIMARY KEY,
  jan BIGINT,
  tag_from_description VARCHAR(255),
  FOREIGN KEY (jan) REFERENCES Products(jan)
);