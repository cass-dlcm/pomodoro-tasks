create table user_auth_rate_limits
(
    user_id int not null,
    ip_addr varchar(45) not null,
    failed_auth_count int not null,
    last_failed_auth datetime not null,
    constraint user_auth_rate_limits_pk
        primary key (user_id, ip_addr)
);

