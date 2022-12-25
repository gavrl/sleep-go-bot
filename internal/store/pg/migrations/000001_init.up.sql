create table sleep_rate
(
    id         serial not null unique,
    username   varchar(100) not null,
    rate       bigint not null,
	calories   integer,
	sleep_time interval SECOND(0),
	date  	   date not null	
);