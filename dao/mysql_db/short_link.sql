use short_link_db;

drop table short_link_tab_00;
drop table short_link_tab_01;
drop table short_link_tab_02;
drop table short_link_tab_03;
drop table short_link_tab_04;
drop table short_link_tab_05;
drop table short_link_tab_06;
drop table short_link_tab_07;
drop table short_link_tab_08;
drop table short_link_tab_09;
drop table short_link_tab_10;
drop table short_link_tab_11;
drop table short_link_tab_12;
drop table short_link_tab_13;
drop table short_link_tab_14;
drop table short_link_tab_15;

create table if not exists short_link_tab_00(
    id bigint primary key auto_increment,
    short_link_type int,
    short_link varchar(64) unique not null,
    raw_link_md5 varchar(64) unique not null,
    raw_link_data varchar(2048) not null,
    expire_timestamp bigint,
    create_timestamp bigint,
    update_timestamp bigint
);

create index expire_timestamp_index on short_link_tab_00 (expire_timestamp);

create table short_link_tab_01 like short_link_tab_00;
create table short_link_tab_02 like short_link_tab_00;
create table short_link_tab_03 like short_link_tab_00;
create table short_link_tab_04 like short_link_tab_00;
create table short_link_tab_05 like short_link_tab_00;
create table short_link_tab_06 like short_link_tab_00;
create table short_link_tab_07 like short_link_tab_00;
create table short_link_tab_08 like short_link_tab_00;
create table short_link_tab_09 like short_link_tab_00;
create table short_link_tab_10 like short_link_tab_00;
create table short_link_tab_11 like short_link_tab_00;
create table short_link_tab_12 like short_link_tab_00;
create table short_link_tab_13 like short_link_tab_00;
create table short_link_tab_14 like short_link_tab_00;
create table short_link_tab_15 like short_link_tab_00;


