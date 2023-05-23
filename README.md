# dans_test

# Framework Libraries
    echo
    go get -u github.com/go-sql-driver/mysql
    go get golang.org/x/crypto/bcrypt

    go mod tidy


# Docker 
### Setup Database
    nama database : dbauth
    username : goauth
    password : THTqAOELuFckJZZaBP7Z
    jenis database : MariaDB

### Db Container
    mkdir dbdata
    docker run -itd --name=dbauth -p 3307:3306 -v $(pwd)/dbdata:/var/lib/mysql --env="MYSQL_ROOT_PASSWORD=root212" mariadb:10.5
    docker exec -it dbauth bash
    mysql -u root -proot212

    create database dbauth;
    use dbauth;
    CREATE USER 'goauth'@'%' IDENTIFIED BY 'THTqAOELuFckJZZaBP7Z';
    GRANT ALL PRIVILEGES ON *.* TO 'goauth'@'%';  
    FLUSH privileges;

    name container : dbauth
    docker container start dbauth
    docker inspect dbauth
    docker container stop dbauth
    docker rm dbauth

### Table user
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
	 ('Absor MU','admin@absormu.id','U21AcnREZWwhdmVyeQ==',1,1,1,'2022-11-19 15:37:52.0','SYSTEM',NULL,NULL,0);
### Table user_contact
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


# POSTMAN

Login

    {
        "email": "absor@gmail.com",
        "password": "@bs0R212"
    }

    curl --location --request POST 'http://localhost:9670/dans_test/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "email": "absor@gmail.com",
        "password": "@bs0R212"
    }
    '

Job List

    curl --location --request GET 'http://localhost:9670/dans_test/job-list?description=python&location=berlin&pagination=true' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFic29yQGdtYWlsLmNvbSIsImV4cCI6MTY2ODg0NDQ0NSwibmFtZSI6IlBULiBBYnNvciBJbmRvbmVzaWEiLCJyb2xlX2lkIjoyLCJ1aWQiOiJjZHM2MnJiaGVkN2E3bmEzc2gzZyIsInVzZXJfY29udGFjdF9pZCI6MywidXNlcl9pZCI6NX0.wKUTz4yYDoAJa46A938gGiMhRzUbrod2Onl_Ksz_O0w' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "email": "absor@gmail.com",
        "password": "@bs0R212"
    }
    '

Job Detail by Id

    curl --location --request GET 'http://localhost:9670/dans_test/job-detail/32bf67e5-4971-47ce-985c-44b6b3860cdb' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFic29yQGdtYWlsLmNvbSIsImV4cCI6MTY2ODg1Mjc4OCwibmFtZSI6IlBULiBBYnNvciBJbmRvbmVzaWEiLCJyb2xlX2lkIjoyLCJ1aWQiOiJjZHM4NDEzaGVkNzQyYmZkOTU1MCIsInVzZXJfY29udGFjdF9pZCI6MywidXNlcl9pZCI6NX0.yzFQjzyB62wiEtimf9JMZB44D3VKXQCDQrPRbm53kis' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "email": "absor@gmail.com",
        "password": "@bs0R212"
    }
    '

# SQL Test

    No. 1
        SELECT 
        C.ID, C.CUST_FIRSTNAME, C.CUST_LASTNAME, 
        A.ACC_NUMBER, SUM(A.ACC_BALANCE) AS BALANCE, 
        FROM CUSTOMER C  
        JOIN ACCOUNT A ON A.ACC_OWNER = C.CUST_ID;

    No. 2 
        SELECT 
        T.TRS_DATE, A.ACC_NUMBER, C.CUST_FIRSTNAME, C.CUST_LASTNAME
        FROM `TRANSACTION` T 
        JOIN CUSTOMER C ON C.CUST_ID = A.ACC_OWNER 
        JOIN ACCOUNT A ON  A.ACC_NUMBER = T.TRS_FROM_ACCOUNT
        WHERE C.CUST_FIRSTNAME = "John Michael"    
        ORDER BY T.TRS_DATE DESC, A.ACC_NUMBER DESC; # tokped
