-- set sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
set sql_mode = 'ONLY_FULL_GROUP_BY,ALLOW_INVALID_DATES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';

drop table if exists `dev_time_t`;
create table if not exists `dev_time_t` (
    id int unsigned auto_increment primary key,
    dev_date date not null default '0000-00-00',
    dev_time time not null default '00:00:00',
    dev_datetime datetime not null default '0000-00-00 00:00:00',
    dev_timestamp timestamp not null default '0000-00-00 00:00:00',
    dev_year year not null default 0000
);
