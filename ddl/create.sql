drop table if exists users;
create table users (
  `id` int not null primary key auto_increment,
  `uuid` varchar(255) not null unique,
  `name` varchar(255),
  `email` varchar(255),
  `password` varchar(255),
  `created_at` datetime not null
);

drop table if exists sessions;
create table sessions (
  `id` int not null primary key auto_increment,
  `uuid` varchar(255) not null unique,
  `email` varchar(255),
  `user_uuid` varchar(255),
  `created_at` datetime not null
);

drop table if exists threads;
create table threads (
  `id` int not null primary key auto_increment,
  `uuid` varchar(255) not null unique,
  `topic` text,
  `user_uuid` varchar(255),
  `created_at` datetime not null
);

drop table if exists posts;
create table posts (
  `id` int not null primary key auto_increment,
  `uuid` varchar(255) not null unique,
  `body` text,
  `user_uuid` varchar(255),
  `thread_uuid` varchar(255),
  `created_at` datetime not null
);
