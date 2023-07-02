## CatPicHub

### Requirements

* PostgreSQL@latest
* Go 1.20 amd64 version
* Gin HTTP web framework (available in Golang
* GORM (ORM library for Golang)



### How to Run

* Set up the database :
```
CREATE DATABASE catpichub;
\c catpichub;
CREATE USER catpichub WITH PASSWORD 'catpichub@1234';
ALTER DATABASE catpichub OWNER TO catpichub;
GRANT ALL PRIVILEGES ON DATABASE catpichub TO catpichub;
```

* Run migrations
```
go build -o goose-custom ./migrations/*.go
./goose-custom up
```


### [Postman Collection](https://api.postman.com/collections/16960428-0e10276e-0960-4ffe-9f93-26d39747cb7b?access_key=PMAT-01GN42TH8Q6NE1CF37TZNTR4ZY)

###



