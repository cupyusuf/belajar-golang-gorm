-- Active: 1741153327066@@127.0.0.1@3306@belajar_golang_gorm
create table sample (
    id varchar(100) not null,
    name varchar(100) not null,
    primary key (id)
) engine = InnoDB;

select * from sample;