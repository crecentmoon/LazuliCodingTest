CREATE TABLE IF NOT EXISTS DescriptionTags (
  tag_id INT PRIMARY KEY,
  jan INT,
  tag_from_description VARCHAR(255),
  FOREIGN KEY (jan) REFERENCES Products(jan)
);