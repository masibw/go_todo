CREATE DATABASE go_todo;

USE go_todo;

create table if not exists users (id INT AUTO_INCREMENT, email VARCHAR(256) NOT NULL, password VARCHAR(256) NOT NULL, created_at DATETIME  DEFAULT CURRENT_TIMESTAMP,last_login DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (id), UNIQUE (email) ) ENGINE = INNODB, CHARACTER SET = utf8mb4, COLLATE utf8mb4_general_ci;

create table if not exists todos (id INT AUTO_INCREMENT,user_id INT NOT NULL,content VARCHAR(512) NOT NULL,done BOOLEAN NOT NULL DEFAULT FALSE, created_at DATETIME NOT NULL  DEFAULT CURRENT_TIMESTAMP,updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (id),FOREIGN KEY fk_user (user_id) REFERENCES users(id)) ENGINE = INNODB, CHARACTER SET = utf8mb4, COLLATE utf8mb4_general_ci;