-- auto-generated definition
CREATE TABLE books
(
    id    int AUTO_INCREMENT
        PRIMARY KEY,
    title varchar(255) CHARSET utf8 NULL,
    cover text                      NULL,
    price decimal(18, 2)            NULL
);
