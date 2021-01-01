DROP TABLE IF EXISTS offers;

CREATE TABLE offers
(
    seller_id int not null,
    offer_id int not null,
    name varchar not null,
    price int not null,
    quantity int not null,
    available boolean not null,
    unique (seller_id, offer_id)
);

INSERT INTO offers
VALUES
(1, 1, 'телевизор', 25000, 50, true),
(1, 2, 'телефон', 13500, 125, true),
(2, 1, 'ноутбук', 100000, 15, true),
(3, 1, 'часы', 30000, 75, true);

-- SELECT * FROM offers;