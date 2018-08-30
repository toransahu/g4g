# MySQL

## Install

```
sudo apt install mysql-server
```

## Run

```
mysql -u root -p
```

if you get :

```
ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/var/run/mysqld/mysqld.sock' (2)
```

then check 

```
services mysql start
```

Default password for root is: **password**

## Help
```
mysql> help;

or

\h
```

### chage root password

```
sudo mysqladmin password <new password>
```

## Database

### Install `Sakila` db prepared by mysql
```
wget http://downloads.mysql.com/docs/sakila-db.tar.gz
tar -xzf sakila-db.tar.gz
cd sakila-db
mysql -u root -p < sakila-schema.sql
mysql -u root -p < sakila-data.sql
```

### To see databases
```
mysql> show databases;
```

### To create a database
```
mysql> create database <db_name>;

Query OK, 1 row affected (0.00 sec)
```

### To use a database

```
mysql> use <db_name>;
```

## Tables

### To see tables
```
show tables;
```