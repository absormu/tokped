 
# Framework Libraries
    echo
    go get -u github.com/go-sql-driver/mysql
    go get golang.org/x/crypto/bcrypt

    git init
    go mod init github.com/absormu/tokped
    go mod tidy


### Table 

    create table user_contact(
        id BIGINT(20) NOT NULL AUTO_INCREMENT,
        name VARCHAR(100) NOT NULL,
        city_id INT(11) NOT NULL,
        telephone VARCHAR(30) NULL,
        address TEXT NULL,
        type tinyint(1) NOT NULL, 
        active tinyint(1) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        created_by VARCHAR(100) NOT NULL,
        modified_at TIMESTAMP NULL, 
        modified_by VARCHAR(100),
        is_deleted tinyint(1) DEFAULT 0 NOT NULL,
        PRIMARY KEY ( id )
    );
    
    create table user(
        id BIGINT(20) NOT NULL AUTO_INCREMENT,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(150) NOT NULL,
        password VARCHAR(150) NOT NULL,
        user_contact_id BIGINT(20) NOT NULL,
        role_id INT(11) NOT NULL,
        active tinyint(1) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        created_by VARCHAR(100) NOT NULL,
        modified_at TIMESTAMP NULL, 
        modified_by VARCHAR(100),
        is_deleted tinyint(1) DEFAULT 0 NOT NULL,
        PRIMARY KEY ( id )
    );

    INSERT INTO `user` (name,email,password,user_contact_id,role_id,active,created_at,created_by,modified_at,modified_by,is_deleted) VALUES
	 ('Absor MU','admin@absormu.id','$2a$12$.ixO2S9J6gcOhIIRdXKZRuRU36F5DMDl6Qi1rw2kfHICnNbcsQaL6',1,1,1,'2022-11-19 15:37:52.0','SYSTEM',NULL,NULL,0);



# POSTMAN

Login

    {
        "email": "admin@absormu.id",
        "password": "@bs0R212"
    }

    curl --location --request POST 'http://localhost:9670/dans_test/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{
         "email": "admin@absormu.id",
        "password": "@bs0R212"
    }
    '

Job List

    curl --location --request GET 'http://localhost:9670/dans_test/job-list?description=python&location=berlin&pagination=true' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFic29yQGdtYWlsLmNvbSIsImV4cCI6MTY2ODg0NDQ0NSwibmFtZSI6IlBULiBBYnNvciBJbmRvbmVzaWEiLCJyb2xlX2lkIjoyLCJ1aWQiOiJjZHM2MnJiaGVkN2E3bmEzc2gzZyIsInVzZXJfY29udGFjdF9pZCI6MywidXNlcl9pZCI6NX0.wKUTz4yYDoAJa46A938gGiMhRzUbrod2Onl_Ksz_O0w' \
    --header 'Content-Type: application/json' \
    --data-raw '{
         "email": "admin@absormu.id",
        "password": "@bs0R212"
    }
    '

Job Detail by Id

    curl --location --request GET 'http://localhost:9670/dans_test/job-detail/32bf67e5-4971-47ce-985c-44b6b3860cdb' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFic29yQGdtYWlsLmNvbSIsImV4cCI6MTY2ODg1Mjc4OCwibmFtZSI6IlBULiBBYnNvciBJbmRvbmVzaWEiLCJyb2xlX2lkIjoyLCJ1aWQiOiJjZHM4NDEzaGVkNzQyYmZkOTU1MCIsInVzZXJfY29udGFjdF9pZCI6MywidXNlcl9pZCI6NX0.yzFQjzyB62wiEtimf9JMZB44D3VKXQCDQrPRbm53kis' \
    --header 'Content-Type: application/json' \
    --data-raw '{
         "email": "admin@absormu.id",
        "password": "@bs0R212""
    }
    '
Order History by Id

	 curl --location 'localhost:9670/tokped/order-history/1' \
	--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFic29yYWx2b3JkMDdAZ21haWwuY29tIiwiZXhwIjoxNjg0ODg5ODUwLCJuYW1lIjoiTXVoYW1hZCBVbGlsIEFic29yIiwicm9sZSI6ImFkbWluIiwicm9sZV9pZCI6MiwidWlkIjoiY2htMXIyaGRrYnRnYnNtb3E5YWciLCJ1c2VyX2NvbnRhY3RfaWQiOjMsInVzZXJfaWQiOjN9.BrTpohZLD3Dr2uzHhIh7zQPSqZBwIuWyurdJ7LSlhPs'
