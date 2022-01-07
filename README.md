# Go
A go assesment

using Golang Gin and a MySQL database:
write a REST API with the following tables:

create table person(id int not null key auto_increment,  name varchar(255), age int);
You can assume the following data in person:
Name, age, id
mike , 31, 1
John, 45, 2
Joseph, 20, 3

create table phone(id int not null key auto_increment,  number varchar(255), person_id int);
You can assume the following data in phone:
person_id, id, number
1,1, 444-444-4444
2,8, 123-444-7777
3,3, 445-444-4444

create table address(id int not null key auto_increment,  city varchar(255), state varchar(255), street1 varchar(255), street2 varchar(255), zip_code varchar(255));

You can assume the following data in address:
id, city, state, street1, street2, zip_code
1,Eugene, OR, "111 Main St", "", 98765
2, Sacramento, CA, "432 First St", "Apt 1", 22221
3, Austin, TX, "213 South 1st St", 78704

 create table address_join(id int not null key auto_increment,  person_id int, address_id int);
You can assume the following data in address_join:
id, person_id, address_id
1,1,3
2,2,1
3,3,2

Task 1:
make a REST endpoint (GET)
/person/{person_id}/info
GET /person/1/info  returns:
{
"name": "mike",
"phone_number": "444-444-4444",
"city" : "Eugene",
"state" : "OR",
"street1": "111 Main St",
"street2": "", 
"zip_code": "98765",
}

Task 2:
make a REST endpoint (POST)  / Create
/person/create
POST /person/create  accepts body: 
{
"name": "mike",
"phone_number": "444-444-4444",
"city" : "Eugene",
"state" : "OR",
"street1": "111 Main St",
"street2": "", 
"zip_code": "98765",
}
returns 200
