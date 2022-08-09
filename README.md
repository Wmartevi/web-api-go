# web-api-go

### Create module
```
$ go mod init
```
### Install libraries
```
$ go install github.com/gin-gonic/gin@latest
```
```
go install gorm.io/gorm@latest
```
```
$ go install gorm.io/driver/postgres@latest
```

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
