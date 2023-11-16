CREATE DATABASE IF NOT EXISTS wallet_core;

USE wallet_core;

CREATE TABLE IF NOT EXISTS payment_type (
    id INT NOT NULL PRIMARY KEY,
    type_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS invoice_category (
    id INT NOT NULL PRIMARY KEY,
    category VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS gain_category (
    id INT NOT NULL PRIMARY KEY,
    category VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS invoice_projection (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    created_at INT NOT NULL,
    buy_at DATE NOT NULL,
    pay_in DATE NOT NULL,
    description VARCHAR(255) NOT NULL,
    value DECIMAL(15,2) NOT NULL,
    is_done BOOLEAN NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    payment_type_id INT NOT NULL,
    category_id INT NOT NULL,
    CONSTRAINT FK_invoice_projection_payment_type FOREIGN KEY (payment_type_id) REFERENCES payment_type(id),
    CONSTRAINT FK_invoice_projection_category FOREIGN KEY (category_id) REFERENCES invoice_category(id)
);

CREATE TABLE IF NOT EXISTS invoice_done (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    created_at INT NOT NULL,
    buy_at DATE NOT NULL,
    pay_at DATE NOT NULL,
    description VARCHAR(255) NOT NULL,
    value DECIMAL(15,2) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    payment_type_id INT NOT NULL,
    category_id INT NOT NULL,
    invoice_projection_id VARCHAR(255),
    CONSTRAINT FK_invoice_done_payment_type FOREIGN KEY (payment_type_id) REFERENCES payment_type(id),
    CONSTRAINT FK_invoice_done_category FOREIGN KEY (category_id) REFERENCES invoice_category(id),
    CONSTRAINT FK_invoice_done_invoice_projection FOREIGN KEY (invoice_projection_id) REFERENCES invoice_projection(id)
);

CREATE TABLE IF NOT EXISTS gain_projection (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    created_at INT NOT NULL,
    pay_in DATE NOT NULL,
    description VARCHAR(255) NOT NULL,
    value DECIMAL(15,2) NOT NULL,
    is_done BOOLEAN NOT NULL,
    is_passive BOOLEAN NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    CONSTRAINT FK_gain_projection_category FOREIGN KEY (category_id) REFERENCES gain_category(id)
);

CREATE TABLE IF NOT EXISTS gain_done (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    created_at INT NOT NULL,
    pay_in DATE NOT NULL,
    description VARCHAR(255) NOT NULL,
    value DECIMAL(15,2) NOT NULL,
    is_passive BOOLEAN NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    gain_projection_id VARCHAR(255),
    CONSTRAINT FK_gain_done_category FOREIGN KEY (category_id) REFERENCES gain_category(id),
    CONSTRAINT FK_gain_done_gain_projection FOREIGN KEY (gain_projection_id) REFERENCES gain_projection(id)
);

CREATE TABLE IF NOT EXISTS label (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    label VARCHAR(255) NOT NULL
);
