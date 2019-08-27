# drop DATABASE db_apiserver;
CREATE DATABASE IF NOT EXISTS db_apiserver CHARACTER SET UTF8MB4;

USE db_apiserver;

DROP TABLE IF EXISTS tb_users;

CREATE TABLE tb_users(
                         id BIGINT(20) UNSIGNED NOT NULL auto_increment,
                         username VARCHAR(255) NOT NULL ,
                         password VARCHAR(255) NOT NULL ,
                         createdAt TIMESTAMP NULL DEFAULT NULL,
                         updatedAt TIMESTAMP NULL DEFAULT NULL,
                         deletedAt TIMESTAMP NULL DEFAULT NULL,
                         PRIMARY KEY (id),
                         UNIQUE KEY username (username),
                         KEY idx_tb_users_deletedAt (deletedAt)
)ENGINE=INNODB AUTO_INCREMENT=0 DEFAULT  CHARSET=UTF8MB4;

LOCK TABLES tb_users WRITE ;
INSERT INTO tb_users VALUEs (0,'admin','123456',DEFAULT,DEFAULT,DEFAULT);
UNLOCK TABLES ;

select * from tb_users;
show create table tb_users;