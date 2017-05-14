# jackmarshall

``` sh
cd auth
vagrant up
vagrant ssh
cd go get ./
make redis
go get ./...
go build
./server
```

``` sh
cd front
npm start
gulp watch
```

``` sh
cd tournaments
docker run -p 27017:27017 mongo
go get ./...
go build
./tournament
```
