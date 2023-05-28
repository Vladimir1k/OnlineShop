

/* Вывести id продавца, имя продавца, количество товара*/
SELECT sellers.id, sellers.name, COUNT(products.name) AS products_count
FROM sellers
         LEFT JOIN products ON sellers.id = products.id_seller
GROUP BY sellers.id;

/* Вывести product_id, product_name, seller_name, seller_phone*/
SELECT products.id, products.name AS product_name, sellers.name AS seller_name, sellers.phone
FROM products
         FULL OUTER JOIN sellers ON products.id_seller = sellers.id;

/*Разные товары от разных продавцов(не уверен правильно ли понял задание, потому что легкий запрос))*/
SELECT DISTINCT products.name AS product_name, sellers.name AS seller_name
FROM products
         INNER JOIN sellers ON products.id_seller = sellers.id;

/*ТОП 10 товаров, которые купили больше всего*/
SELECT p.id AS product_id, p.name AS product_name, s.name AS seller_name,
       COUNT(o.product_id) AS purchased_quantity
FROM products p
         INNER JOIN orders o ON p.id = o.product_id
         INNER JOIN sellers s ON p.id_seller = s.id
GROUP BY p.id, s.name
ORDER BY purchased_quantity DESC
    LIMIT 10;

/*Все покупатели которые потратили больше 500*/
SELECT b.id, b.name, SUM(p.price * o.quantity) AS orders_sum
FROM buyers b
         INNER JOIN orders o ON b.id = o.buyer_id
         INNER JOIN products p ON o.product_id = p.id
GROUP BY b.id
HAVING SUM(p.price * o.quantity) > 500
ORDER BY orders_sum DESC;


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



