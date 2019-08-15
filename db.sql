CREATE TABLE IF NOT EXISTS project_brief(
	id int(32),
	seq varchar(10),
	pre_number varchar(50),
	number_link varchar(100),
	name varchar(50),
	name_link varchar(100),
	company varchar(50),
	district varchar(50),
	approve_time varchar(30),
	primary key(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;