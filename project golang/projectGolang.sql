create database Cloud;
use Cloud;
create table Users
(
    ID int auto_increment primary key,
    UserName varchar(50),
    Pass varchar(200) ,
	Email varchar(255) ,
    PhoneNumber int null,
    Wallet float,
    Credit float,
    Address varchar(255),
    VIPuser varchar(100)
);
create table admins
(
	ID int auto_increment primary key,
    UserName varchar(50),
    Pass varchar(200) ,
    Email varchar(255) ,
    PhoneNumber varchar(50) null,
    Address varchar(255) null,
    VIPadmins varchar(100)
);
create table products
(
	ID int auto_increment primary key,
    NameProduct varchar(255),
    Descriptions varchar(255),
    parent int
    
);

create table productsPackage
(
	ID int auto_increment primary key,
    NameProduct varchar(255),
    RAM varchar(255) null,
    CPU varchar(255) null,
    Storage varchar(255) null,
    Price float,
    ProductID int,
    foreign key (ProductID) references products (ID)
    
);