SHOW DATABASES;

-- drop database children_math;

CREATE DATABASE children_math;

USE children_math;

CREATE TABLE users(
	user_name VARCHAR(20) PRIMARY KEY,
	`password` VARCHAR(50),
	points INT
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_bin;

ALTER TABLE users
ADD INDEX idx_users_points_username(points DESC, user_name ASC);

-- DROP TABLE problems;

CREATE TABLE problems(
	id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
	user_name VARCHAR(20),
	num1 INT,
	num2 INT,
	wrong_ans INT,
	operator CHAR(1),
	CONSTRAINT fk_problems_username
	FOREIGN KEY(user_name)
	REFERENCES users(user_name)
)ENGINE=INNODB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_bin;

ALTER TABLE problems
ADD INDEX idx_problems_username(user_name);
