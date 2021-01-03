CREATE TABLE offers (
    seller_id int NOT NULL CHECK (seller_id > 0),
    offer_id int NOT NULL CHECK (offer_id > 0),
    name varchar NOT NULL,
    price int NOT NULL CHECK (price > 0),
    quantity int NOT NULL CHECK (price > 0),
    available boolean NOT NULL,
    UNIQUE (seller_id, offer_id)
);

CREATE TABLE stats (
    id serial PRIMARY KEY,
    status varchar NOT NULL,
    create_count int NOT NULL CHECK (create_count >= 0),
    update_count int NOT NULL CHECK (update_count >= 0),
    delete_count int NOT NULL CHECK (delete_count >= 0),
    error_count int NOT NULL CHECK (error_count >= 0)
);