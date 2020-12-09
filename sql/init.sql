
CREATE TABLE IF NOT EXISTS urls(
  id int NOT NULL AUTO_INCREMENT,
  url varchar(4000) NOT NULL,
  short_url varchar(45) NOT NULL,
  create_at timestamp default current_timestamp NOT NULL,
  update_at timestamp default current_timestamp NOT NULL,
  delete_at datetime DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY short_url_UNIQUE (short_url)
) 
ENGINE=InnoDB 
AUTO_INCREMENT=16 
DEFAULT CHARSET=utf8mb4 
COLLATE=utf8mb4_0900_ai_ci;