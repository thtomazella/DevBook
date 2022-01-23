CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

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
    FOREIGN KEY (seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE,

    primary key (usuario_id, seguidor_id)

 )ENGINE=INNODB;

CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo VARCHAR(50) not null,
    conteudo VARCHAR(300) not null,

    autor_id int not null,
    FOREIGN KEY(autor_id) REFERENCES usuarios(id)  ON DELETE CASCADE,

    curtidas int DEFAULT 0,
    criadoEm timestamp DEFAULT current_timestamp
) ENGINE=INNODB;)
