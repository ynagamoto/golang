CREATE DATABASE test_login;
USE test_login;
CREATE TABLE users (id int not null auto_increment primary key, name varchar(32) not null, passwd varchar(64) not null);
