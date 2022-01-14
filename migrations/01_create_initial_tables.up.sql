CREATE TABLE some_data
(
    symbol varchar(255) not null unique,
    price double precision not null,
    volume double precision not null,
    last_trade double precision not null
);