CREATE DATABASE proj_db;

USE proj_db;

CREATE TABLE Customer (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    email VARCHAR(50) NOT NULL,
    username VARCHAR(45) NOT NULL,
    `password` VARCHAR(254) NOT NULL,
    PRIMARY KEY(id),
    UNIQUE(email, username)
);

CREATE TABLE BOwner (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    email VARCHAR(50) NOT NULL,
    `password` VARCHAR(254) NOT NULL,
    contact VARCHAR(15) NOT NULL,
    PRIMARY KEY(id),
    UNIQUE(email)
);

CREATE TABLE Category (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(20) NOT NULL,
    PRIMARY KEY(id),
    UNIQUE(title)
);

CREATE TABLE Listing (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    shop_title VARCHAR(45) NOT NULL,
    shop_description TEXT NOT NULL,
    ig_url VARCHAR(100),
    fb_url VARCHAR(100),
    website_url VARCHAR(100),
    bowner_id INT UNSIGNED NOT NULL,
    category_id INT UNSIGNED NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(bowner_id) REFERENCES BOwner(id),
    FOREIGN KEY(category_id) REFERENCES Category(id)
);

CREATE TABLE FoodItem (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    item_name VARCHAR(25) NOT NULL,
    price INT UNSIGNED NOT NULL,
    description VARCHAR(254),
    listing_id INT UNSIGNED NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(listing_id) REFERENCES Listing(id)
);

CREATE TABLE Review (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    review TEXT NOT NULL,
    customer_id INT UNSIGNED NOT NULL,
    listing_id INT UNSIGNED NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(customer_id) REFERENCES Customer(id),
    FOREIGN KEY(listing_id) REFERENCES Listing(id)
);


INSERT INTO Category (title) VALUES ('Dim Sum');
