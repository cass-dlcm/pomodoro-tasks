create table lists
(
    id   int auto_increment,
    listName varchar(63) not null,
    constraint lists_id_uindex
        unique (id)
);

alter table lists
    add primary key (id);

