# jackmarshall/auth

``` sh
vagrant up
vagrant ssh
cd src/github.com/chibimi/jackmarshall/auth/server/
make redis
go get ./...
go build
./server
```
