# go-simple-bulk-insert

> this project is developed with **clean architechture** environment. \
> (at least, it's clean enough for me)

## development requirements:
- go version go1.23.4 linux/amd64
- go Fiber v2.52.6
- mysql  Ver 8.0.41-0ubuntu0.24.04.1 for Linux on x86_64 ((Ubuntu))

## Simple installations:
- run **tables.sql** file in **docs/** to your mysql database
- run `go mod init go-simple-bulk-insert`
- run `go mod tidy`
- run `go run main.go` for quick testing
- run `go build` for node testing