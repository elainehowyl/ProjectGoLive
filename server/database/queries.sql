CREATE DATABASE proj_db;

USE proj_db;

CREATE TABLE BOwner (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    email VARCHAR(50) NOT NULL,
    `password` VARCHAR(254) NOT NULL,
    contact VARCHAR(15) NOT NULL,
    PRIMARY KEY(id),
    UNIQUE(email)
);

CREATE TABLE BOwner (
    email VARCHAR(50) NOT NULL,
    `password` VARCHAR(254) NOT NULL,
    contact VARCHAR(15) NOT NULL,
    PRIMARY KEY(email)
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

CREATE TABLE Listing (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    shop_title VARCHAR(45) NOT NULL,
    shop_description TEXT NOT NULL,
    ig_url VARCHAR(100),
    fb_url VARCHAR(100),
    website_url VARCHAR(100),
    bowner_email VARCHAR(50) NOT NULL,
    category_id INT UNSIGNED NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(bowner_email) REFERENCES BOwner(email),
    FOREIGN KEY(category_id) REFERENCES Category(id)
);

CREATE TABLE Item (
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
    `name` VARCHAR(45) NOT NULL,
    review TEXT NOT NULL,
    listing_id INT UNSIGNED NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(listing_id) REFERENCES Listing(id)
);

INSERT INTO Category (title) VALUES ('Dim Sum');

INSERT INTO Listing (shop_title, shop_description, ig_url, fb_url, website_url, bowner_id, category_id) 
VALUES ('Traditional HK Cafe', 'A traditional HK cafe opened since the 1950s, well known for delicious egg tarts.', 'www.instagram.com/hkcafe','www.facebook.com/hkcafe','www.hkcafe.com',4,1);

SELECT BOwner.id, email, contact, Listing.id, shop_title, shop_description, ig_url, fb_url, website_url, Category.id, title
FROM BOwner 
JOIN Listing ON BOwner.id=Listing.bowner_id
JOIN Category ON Category.id=Listing.category_id
WHERE email='BOwner@example.com';


SELECT Listing.id, shop_title, shop_description, ig_url, fb_url, website_url, Category.id, title
FROM Listing
JOIN Category
ON Listing.category_id=Category.id
WHERE bowner_id=4\G