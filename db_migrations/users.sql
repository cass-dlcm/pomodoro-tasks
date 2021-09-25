create table users
(
    id       int unsigned auto_increment,
    username varchar(15)                  not null,
    password char(60) collate utf8mb4_bin not null,
    constraint id_UNIQUE
        unique (id),
    constraint username_UNIQUE
        unique (username)
);

alter table users
    add primary key (id);

