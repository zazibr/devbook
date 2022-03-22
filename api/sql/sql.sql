CREATE DATABSE IF NOT EXISTS devbook;
use devbook;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    ID int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(20) not null unique,
    criadoEM timeStamp default current_timestamp()
)ENGINE=INNODB;

-- desc usuarios;