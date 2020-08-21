-- auto-generated definition
CREATE TABLE pos_terminals
(
    id       int AUTO_INCREMENT
        PRIMARY KEY,
    store_id int          NOT NULL,
    username varchar(255) NULL,
    password binary(60)   NULL,
    CONSTRAINT pos_terminal_stores_id_fk
        FOREIGN KEY (store_id) REFERENCES stores (id)
);

