# CatPicHub

## Requirements

* PostgreSQL@latest
* Go 1.20 amd64 version
* Gin HTTP web framework (available in Golang
* GORM (ORM library for Golang)



## How to Run

### Running locally
* Set up the database :
```
CREATE DATABASE catpichub;
\c catpichub;
CREATE USER catpichub WITH PASSWORD 'catpichub@1234';
ALTER DATABASE catpichub OWNER TO catpichub;
GRANT ALL PRIVILEGES ON DATABASE catpichub TO catpichub;
```

* Paste these configs in ./config/env.dev.toml
```
[application]
    app_name = "reminders"
    listen_port = 8080
    listen_ip = "127.0.0.1"

[database]
    dialect = "postgres"
    host = "localhost"
    port = 5432
    protocol = "tcp"
    name = "catpichub"
```

* Run migrations
```
go build -o goose-custom ./migrations/*.go
./goose-custom up
```
* Start the service
```
go build -o catpichub ./cmd/main.go
./catpichub
```

### Running via Docker

* Paste these configs in ./config/env.dev.toml
```
[application]
    app_name = "reminders"
    listen_port = 8080
    listen_ip = "0.0.0.0"

[database]
    dialect = "postgres"
    host = "postgres"
    port = 5432
    protocol = "tcp"
    name = "catpichub"
```
* Run following commands
```
 docker-compose build 
 docker-compose up
```

### [Postman Collection](https://api.postman.com/collections/27596999-4763bc11-1cb5-4166-a91b-8abfc8bf8ba8?access_key=PMAT-01H4BVR33NKXM8DNG34ZEA4A38)

### [API Documentation](https://documenter.getpostman.com/view/27596999/2s93zCYfPw#aa26aaff-3be9-4e13-9462-da6f6eb7b492)



