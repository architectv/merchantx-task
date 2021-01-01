CREATE TABLE offers (
    seller_id int NOT NULL,
    offer_id int NOT NULL,
    name varchar NOT NULL,
    price int NOT NULL,
    quantity int NOT NULL,
    available boolean NOT NULL,
    UNIQUE (seller_id, offer_id)
);