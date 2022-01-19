# data-base-connection
A simple example of how to connect with a MySQL local database using Golang

## Dependencies

You will need to install Golang in your environment.

You also need to install MySQL in your environment and configure a user (see `stringConnection`) in the code.

Also, create a database called `godatabase` with a table called `users`

An external package is required:
```
$ github.com/go-sql-driver/mysql
```

### How to run

```
$ go run main.go
Open connection
&{0xc0000ea000 0x4b8ca0 0xc0000b60f0 <nil> <nil> {{0 0} 0 0 0 0} false <nil> []}
```
