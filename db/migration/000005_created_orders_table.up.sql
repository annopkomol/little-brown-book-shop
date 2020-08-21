-- auto-generated definition
CREATE TABLE orders
(
    id      int AUTO_INCREMENT
        PRIMARY KEY,
    cart_id int NOT NULL,
    book_id int NOT NULL,
    qty     int NOT NULL,
    CONSTRAINT orders_books_id_fk
        FOREIGN KEY (book_id) REFERENCES books (id),
    CONSTRAINT orders_carts_cart_id_fk
        FOREIGN KEY (cart_id) REFERENCES carts (id)
);

