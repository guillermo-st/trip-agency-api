CREATE TABLE IF NOT EXISTS roles(
role_id serial PRIMARY KEY,
role_name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users(
user_id serial PRIMARY KEY,
first_name VARCHAR(100) NOT NULL,
last_name VARCHAR(100) NOT NULL,
email VARCHAR(100) UNIQUE NOT NULL,
password VARCHAR(100) NOT NULL,
role_id INT NOT NULL,
FOREIGN KEY (role_id)
	REFERENCES roles (role_id)
);

CREATE TABLE IF NOT EXISTS trips(
trip_id serial PRIMARY KEY,
user_id INT NOT NULL,
started_at TIMESTAMP NOT NULL,
finished_at TIMESTAMP,
FOREIGN KEY (user_id)
	REFERENCES users (user_id)
);

INSERT INTO roles (role_name) VALUES ('admin'), ('driver');

-- This is an example root user with admin privileges, for dev and testing only. 
-- The password for this user is 'complicatedrootpassword', and it's stored as a hash produced by 10 rounds of Bcrypt 
INSERT INTO users (first_name, last_name, email, password, role_id)
VALUES ('Roothie', 'Roothers', 'root@tripagency.com', '$2a$10$jI86yiZe62DzqqImw/H/WubezL6SXwQM3Oy0bjN0QOS.Ns54vLd2G', 1 )
