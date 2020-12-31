DROP TABLE IF EXISTS offers;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS sellers;

CREATE TABLE sellers
(
    id serial primary key
);

CREATE TABLE products
(
    id serial primary key,
    name varchar
);

CREATE TABLE offers
(
    product_id int references products (id) on delete cascade,
    seller_id int references sellers (id) on delete cascade,
    offer_id int not null,
    price int not null,
    quantity int not null,
    available boolean not null,
    unique (seller_id, offer_id)
);

INSERT INTO sellers
VALUES (1), (2), (3);

INSERT INTO products (name)
VALUES ('телевизор'), ('телефон'), ('ноутбук'), ('часы');

INSERT INTO offers
VALUES
(1, 1, 1, 25000, 50, true),
(2, 1, 2, 13500, 125, true),
(3, 2, 1, 100000, 15, true),
(4, 3, 1, 30000, 75, true);

-- SELECT * FROM sellers;
-- SELECT * FROM products;
-- SELECT * FROM offers;