
CREATE TABLE IF NOT EXISTS blogs (
                                        id INT PRIMARY KEY AUTO_INCREMENT,
                                        title VARCHAR(255) NOT NULL,
                                        content TEXT,
                                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
