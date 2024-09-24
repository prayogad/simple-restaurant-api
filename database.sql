-- Active: 1707988890724@@localhost@5433@simple_resto_web@public
CREATE TABLE makanan_minuman (
    id SERIAL NOT NULL,
    nama VARCHAR(100) NOT NULL,
    harga NUMERIC(10, 2) NOT NULL,
    stok INT NOT NULL
);

ALTER TABLE makanan_minuman
    ADD PRIMARY KEY (id);


CREATE TABLE customer (
    id SERIAL NOT NULL,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT unique_username UNIQUE (username)
);

CREATE TABLE orders (
    id SERIAL NOT NULL,
    quantity INT NOT NULL,
    total_harga NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE orders
    ADD CONSTRAINT fk_order_customer FOREIGN KEY (id_customer) REFERENCES customer (id);

ALTER TABLE orders
    ADD PRIMARY KEY (id);

CREATE TABLE order_detail (
    id_order INT NOT NULL,
    id_makanan INT NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (id_order, id_makanan)
);

ALTER TABLE order_detail
    ADD CONSTRAINT fk_order_detail_order FOREIGN KEY (id_order) REFERENCES orders (id);

ALTER TABLE order_detail
    ADD CONSTRAINT fk_order_detail_makanan FOREIGN KEY (id_makanan) REFERENCES makanan_minuman (id);

SELECT * FROM customer;

ALTER TABLE customer
    ADD COLUMN token VARCHAR(100);