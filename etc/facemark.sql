create table if not exists user
(
	user_id int primary key not null  auto_increment,
	user_name varchar(64),
	user_pass varchar(64),
	user_mobile varchar(32),
	user_type enum('user', 'admin', 'test') not null,
	add_time timestamp not null default CURRENT_TIMESTAMP
);

insert into user (user_name, user_pass) values('admin', 'admin');



