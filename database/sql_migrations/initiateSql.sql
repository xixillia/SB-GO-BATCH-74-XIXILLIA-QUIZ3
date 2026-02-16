-- +migrate Up

CREATE TABLE Categories(
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    created_at TIMESTAMP,
    created_by VARCHAR(256),
    modified_at TIMESTAMP,
    modified_by VARCHAR(256)
);

CREATE TABLE Books(
    id SERIAL PRIMARY KEY,
    title VARCHAR(256),
    description VARCHAR(256),
    image_url VARCHAR(256),
    release_year INT,
    price INT,
    total_page INT,
    thickness VARCHAR(256),
    created_at TIMESTAMP,
    created_by VARCHAR(256),
    modified_at TIMESTAMP,
    modified_by VARCHAR(256),
    category_id INT REFERENCES Categories(id)
);

CREATE TABLE Users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(256),
    password VARCHAR(256),
    created_at TIMESTAMP,
    created_by VARCHAR(256),
    modified_at TIMESTAMP,
    modified_by VARCHAR(256)
);