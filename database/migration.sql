CREATE TABLE foods 
( 
    id int auto_increment, 
    name nvarchar(20) not null, 
    country nvarchar(20), 
    city nvarchar(20), 
    vegeterian int, 
    primary key(id)
);