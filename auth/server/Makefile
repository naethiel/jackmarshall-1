dependencies:
	go get ./...

binary: dependencies *.go
	go build -o authServer

binary-static:
	go build -o authServer -ldflags "-s -linkmode external -extldflags -static -w" -tags docker

pkg: dependencies binary-static
	docker build -t jackmarshall/auth .
	docker save jackmarshall/auth > jackmarshall-auth.tar

clean:
	rm -f jackmarshall-auth.tar auth

docker-run: pkg
	docker run -d --name redis redis
	docker run -p 8081:80 -it --link redis:redis  --name jackmarshallA jackmarshall/auth

docker-rm:
	docker rm --force redis
	docker rm --force jackmarshallA

redis:
	#password = password
	redis-cli FLUSHALL
	redis-cli SET user:1 '{"id":1,"login":"Emilie","password":"$$2a$$12$$2sL26WwNWZ2vIXmJG86O9ulhOHXnzpHslow97Pi3F1fC7jqCfR5MC","roles":["root"]}'
	redis-cli SET user:2 '{"id":2,"login":"Conelle","password":"$$2a$$12$$2sL26WwNWZ2vIXmJG86O9ulhOHXnzpHslow97Pi3F1fC7jqCfR5MC","roles":["organizer"]}'
	redis-cli SET user:3 '{"id":3,"login":"Camille","password":"$$2a$$12$$2sL26WwNWZ2vIXmJG86O9ulhOHXnzpHslow97Pi3F1fC7jqCfR5MC","roles":null}'
	redis-cli HSET users "Emilie" 1
	redis-cli HSET users "Conelle" 2
	redis-cli HSET users "Camille" 3
	redis-cli INCR users_max
	redis-cli INCR users_max
	redis-cli INCR users_max
