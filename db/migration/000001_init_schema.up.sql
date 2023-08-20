CREATE TABLE users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  occupation VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  avatar_file_name VARCHAR(255) NOT NULL,
  role VARCHAR(255) NOT NULL,
  token VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE campaigns (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  name VARCHAR(255),
  short_description VARCHAR(255),
  description TEXT,
  goal_amount INT,
  current_amount INT,
  perks TEXT,
  becker_count INT,
  slug VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE campaign_images (
  id INT PRIMARY KEY AUTO_INCREMENT,
  campaign_id INT NOT NULL,
  file_name VARCHAR(255),
  is_primary TINYINT(1),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
  id INT PRIMARY KEY AUTO_INCREMENT,
  campaign_id INT NOT NULL,
  user_id INT NOT NULL,
  amount INT,
  status VARCHAR(255),
  code VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_name ON users (name);
CREATE INDEX idx_campaigns_user_id ON campaigns (user_id);
CREATE INDEX idx_campaign_images_campaign_id ON campaign_images (campaign_id);
CREATE INDEX idx_transactions_campaign_id ON transactions (campaign_id);
CREATE INDEX idx_transactions_user_id ON transactions (user_id);

ALTER TABLE campaigns ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE campaign_images ADD FOREIGN KEY (campaign_id) REFERENCES campaigns (id);
ALTER TABLE transactions ADD FOREIGN KEY (campaign_id) REFERENCES campaigns (id);
ALTER TABLE transactions ADD FOREIGN KEY (user_id) REFERENCES users (id);
