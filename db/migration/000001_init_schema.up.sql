CREATE TABLE `customer` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `name`                  VARCHAR(50)     NOT NULL,
    `table_no`              INT             NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `orders` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `customer_id`           BIGINT          NOT NULL,
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_order_customer` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`)
);

CREATE TABLE `menu` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `name`                  VARCHAR(50)     NOT NULL,
    `stock`                 INT             NOT NULL,
    `price`                 INT             NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `detail` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `order_id`              BIGINT          NOT NULL,
    `menu_id`               BIGINT          NOT NULL,
    `amount`                BIGINT          NOT NULL,
    `total`                 BIGINT          NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_detail_order` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
    CONSTRAINT `fk_detail_menu` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`)
);



