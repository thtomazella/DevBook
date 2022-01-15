CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id       INT         auto_increment primary key,
    nome     VARCHAR(50) NOT NULL, 
    nick     VARCHAR(50) NOT NULL UNIQUE, 
    email    VARCHAR(50) NOT NULL UNIQUE, 
    senha    VARCHAR(20) NOT NULL UNIQUE,
    criadoem TIMESTAMP   DEFAULT current_timestamp()
 )ENGINE=INNODB;
    

