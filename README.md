# web-api-go

### create module

- go mod init
### install libraries

- go get github.com/gin-gonic/gin
- go get gorm.io/gorm
- go get gorm.io/driver/postgres

### API Archives

- DML BD file at archives/DML_banco.sql

### Run Database

```
$ docker-compose up --build
```

### Run Database

```
$ go run main.go
```

### Migrations

- Migrations is auto runned when start the database