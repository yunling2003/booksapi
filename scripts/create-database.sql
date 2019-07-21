-- in bash window, run below command to create linux user
-- adduser youxintea

-- su - youxintea
-- psql

-- then run below command in psql


create database bookstore with encoding 'UTF-8';

grant all privileges on database bookstore to youxin-teaapi;