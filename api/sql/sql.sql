CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;

CREATE TABLE usuarios(
    id       INT         auto_increment primary key,
    nome     VARCHAR(50) NOT NULL, 
    nick     VARCHAR(50) NOT NULL unique, 
    email    VARCHAR(50) NOT NULL unique, 
    senha    VARCHAR(100) NOT NULL ,
    criadoem TIMESTAMP   DEFAULT current_timestamp()
 )ENGINE=INNODB;
    
CREATE TABLE seguidores(
    usuario_id   int NOT NULL, 
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id  int NOT NULL,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    primary key (usuario_id, seguidor_id)

 )ENGINE=INNODB;

