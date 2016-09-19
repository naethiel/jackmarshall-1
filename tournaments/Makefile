dependencies:
	go get ./...

binary:
	go build -o tournaments

binary-static: dependencies *.go
	go build -o tournaments -ldflags "-s -linkmode external -extldflags -static -w" -tags docker

pkg: dependencies binary-static
	docker build -t jackmarshall/tournaments .
	docker save jackmarshall/tournaments > jackmarshall-tournaments.tar

docker-run: pkg
	docker run -d --name mongo mongo:3.0
	docker run -p 8080:8080 -it --link mongo:database  --name jackmarshall jackmarshall/tournaments

docker-rm:
	docker rm --force mongo
	docker rm --force jackmarshall
