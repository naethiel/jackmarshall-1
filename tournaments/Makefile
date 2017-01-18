export GOOS=linux

dependencies:
	go get ./...

binary: dependencies *.go
	go build -o tournaments

binary-static: dependencies *.go
	go build -o tournaments -tags docker

pkg: dependencies binary-static
	docker build -t jackmarshall/tournaments .
	docker save jackmarshall/tournaments > jackmarshall-tournaments.tar

docker-run: pkg
	docker run -d --name mongo mongo:3.0
	sleep 15s
	docker run -p 8080:8080 -it --link mongo:database  --name jackmarshallT jackmarshall/tournaments

docker-rm:
	docker rm --force mongo
	docker rm --force jackmarshallT

db-reset-tournaments:
	mongo jackmarshall data/scenario.js
