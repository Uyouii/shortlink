use short_link_db;

drop table short_link_tab_00000000;
drop table short_link_tab_00000001;
drop table short_link_tab_00000002;
drop table short_link_tab_00000003;
drop table short_link_tab_00000004;
drop table short_link_tab_00000005;
drop table short_link_tab_00000006;
drop table short_link_tab_00000007;
drop table short_link_tab_00000008;
drop table short_link_tab_00000009;
drop table short_link_tab_00000010;
drop table short_link_tab_00000011;
drop table short_link_tab_00000012;
drop table short_link_tab_00000013;
drop table short_link_tab_00000014;
drop table short_link_tab_00000015;

create table if not exists short_link_tab_00000000(
    id bigint primary key auto_increment,
    short_link_type int,
    short_link_path varchar(64),
    raw_link_key varchar(64) unique not null,
    raw_link varchar(2048) not null,
    expire_timestamp bigint,
    create_timestamp bigint,
    update_timestamp bigint
);

create index expire_timestamp_index on short_link_tab_00000000 (expire_timestamp);

create table short_link_tab_00000001 like short_link_tab_00000000;
create table short_link_tab_00000002 like short_link_tab_00000000;
create table short_link_tab_00000003 like short_link_tab_00000000;
create table short_link_tab_00000004 like short_link_tab_00000000;
create table short_link_tab_00000005 like short_link_tab_00000000;
create table short_link_tab_00000006 like short_link_tab_00000000;
create table short_link_tab_00000007 like short_link_tab_00000000;
create table short_link_tab_00000008 like short_link_tab_00000000;
create table short_link_tab_00000009 like short_link_tab_00000000;
create table short_link_tab_00000010 like short_link_tab_00000000;
create table short_link_tab_00000011 like short_link_tab_00000000;
create table short_link_tab_00000012 like short_link_tab_00000000;
create table short_link_tab_00000013 like short_link_tab_00000000;
create table short_link_tab_00000014 like short_link_tab_00000000;
create table short_link_tab_00000015 like short_link_tab_00000000;