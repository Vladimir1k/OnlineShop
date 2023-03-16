create database postgres

create table buyers
(
    id integer generated always as identity
    constraint buyer_pkey
    primary key,
    name varchar(255),
    phone varchar(26),
    address varchar(255)
);

create table sellers
(
    id integer generated always as identity
    constraint seller_pkey
    primary key,
    name varchar(255),
    phone varchar(26)
);

create table products2
(
    id integer generated always as identity
    constraint product_pkey
    primary key,
    name varchar,
    description varchar,
    price double precision,
    id_seller integer
    constraint product_id_seller_fkey
    references sellers
);

create table orders
(
    id integer generated always as identity
    constraint shop_order_pkey
    primary key,
    buyer_id integer
    constraint orders_bayer_id_fkey
    references buyers,
    product_id integer
    constraint orders_product_id_fkey
    references products,
    quantity integer
);