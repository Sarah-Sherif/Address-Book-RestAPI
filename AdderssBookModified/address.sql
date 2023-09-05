Create Table Address(
BuildingID nvarchar(10) primary key,
PNO nvarchar(20),
StreetName varchar(20),
StreetNo varchar(10),
ZipCode varchar(10),
City varchar(15),
State varchar(15),
foreign key (PNO) references people(PhoneNumber) 
on Delete restrict
on Update cascade
);