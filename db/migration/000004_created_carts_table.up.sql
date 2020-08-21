-- auto-generated definition
CREATE TABLE carts
(
    id              int AUTO_INCREMENT
        PRIMARY KEY,
    pos_terminal_id int NULL,
    CONSTRAINT carts_pos_terminals_id_fk
        FOREIGN KEY (pos_terminal_id) REFERENCES pos_terminals (id)
);
