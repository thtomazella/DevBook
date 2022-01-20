insert into usuarios(nome, nick, email, senha)
VALUES
("Usuario 1","Usuario_1", "usuario1@gmail.com","$2a$10$NjVVF.ohPIMKJ9DdtJt4Eec9HibdvzNtU9RNODdE4GYL536iChepq"),
("Usuario 2","Usuario_2", "usuario2@gmail.com","$2a$10$NjVVF.ohPIMKJ9DdtJt4Eec9HibdvzNtU9RNODdE4GYL536iChepq"),
("Usuario 3","Usuario_3", "usuario3@gmail.com","$2a$10$NjVVF.ohPIMKJ9DdtJt4Eec9HibdvzNtU9RNODdE4GYL536iChepq");

insert into seguidores(usuario_id, seguidor_id)
VALUES
(1,2),
(3,1),
(1,3);