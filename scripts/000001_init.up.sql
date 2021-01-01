CREATE TABLE offers (
    seller_id int NOT NULL CHECK (seller_id > 0),
    offer_id int NOT NULL CHECK (offer_id > 0),
    name varchar NOT NULL,
    price int NOT NULL CHECK (price > 0),
    quantity int NOT NULL CHECK (price > 0),
    available boolean NOT NULL,
    UNIQUE (seller_id, offer_id)
);