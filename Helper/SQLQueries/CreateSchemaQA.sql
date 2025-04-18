CREATE database URLShortner_QA;
use URLShortner_QA;

Create table URLDATA(
	ID int NOT NULL AUTO_INCREMENT,
    Original_URL varchar(255),
     New_URL varchar(255),
     PRIMARY KEY (ID),
     unique(ID)	           
);