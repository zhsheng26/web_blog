create database web_blog;
create table if not exists posts
(
  id      bigint primary key auto_increment,
  title   varchar(20),
  content varchar(500)
) DEFAULT CHARSET = utf8mb4;
